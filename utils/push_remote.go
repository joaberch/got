package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func PushRemote(ip, user, path, identityPath string) error {
	cmd := exec.Command("scp" /*"-i", identityPath,*/, "-r", ".got/objects", fmt.Sprintf("%s@%s:%s", user, ip, path))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push remote objects: %w", err)
	}

	for _, file := range []string{"commits.csv", "head"} { //The other files than objects to copy
		src := filepath.Join(".got", file)
		dst := fmt.Sprintf("%s@%s:%s/%s", user, ip, path, file)
		cmd = exec.Command("scp", src, dst)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to push remote objects: %w", err)
		}
	}
	fmt.Printf("pushed remote objects: %s\n", path)
	return nil
}
