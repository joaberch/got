package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// GenerateSSHKey generates a 4096 rsa key in the path given
func GenerateSSHKey(identityPath string) error {
	if _, err := os.Stat(identityPath); os.IsNotExist(err) {
		cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-f", identityPath, "-N", "")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return fmt.Errorf("SSH key already exists: %s", identityPath)
}
