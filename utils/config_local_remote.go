package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigLocalRemote adds the path to the .gotconfig file in the local type
func ConfigLocalRemote(path string) error {
	remotePath := filepath.Join(".got", ".gotconfig")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("config Local Remote Path %s not found: %w", path, err)
	}

	remote := map[string]string{"local": path}
	data, err := json.MarshalIndent(remote, "", "  ")
	if err != nil {
		return fmt.Errorf("config Local Remote Marshal failed: %w", err)
	}

	err = os.WriteFile(remotePath, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("config Local Remote WriteFile failed: %w", err)
	}

	fmt.Printf("Config Local Remote WriteFile success: %s\n", path)
	return nil
}
