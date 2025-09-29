package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigRemotePush configurates the connection to the remote host and setup the ssh key par passwordless communication
func ConfigRemotePush(ip, user, path string) error {
	configPath := filepath.Join(".got", ".gotconfig")
	identity := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_rsa")
	identityPub := identity + ".pub"

	//Generate key if not exist
	if err := GenerateSSHKey(identity); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	if err := CopySSHKeyToRemote(identityPub, user, ip); err != nil {
		return fmt.Errorf("failed to copy SSH key: %w", err)
	}

	remote := map[string]map[string]string{
		"remote": {
			"ip":       ip,
			"user":     user,
			"path":     path,
			"identity": identity,
		},
	}

	data, err := json.MarshalIndent(remote, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal remote config: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write remote config: %w", err)
	}

	fmt.Printf("Remote config saved: %s@%s:%s in %s\n", user, ip, path, configPath)
	return nil
}
