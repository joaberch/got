package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
)

// Push the objects to the local path
func Push() error {
	data, err := os.ReadFile(".got/.gotconfig")
	if err != nil {
		return fmt.Errorf("error reading .gotconfig: %w", err)
	}

	var config map[string]any
	if err = json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("error unmarshaling .gotconfig: %w", err)
	}

	if remote, ok := config["remote"].(map[string]any); ok {
		ip := remote["ip"].(string)
		user := remote["user"].(string)
		path := remote["path"].(string)
		identityPath := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_rsa")
		return utils.PushRemote(ip, user, path, identityPath)
	}

	if localPath, ok := config["local"].(string); ok {
		return utils.PushLocal(localPath)
	}

	return fmt.Errorf("no remote or local config found")
}
