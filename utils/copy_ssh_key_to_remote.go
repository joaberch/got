package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// CopySSHKeyToRemote copy the ssh key to the remote host
func CopySSHKeyToRemote(identityPubPath, user, ip string) error {
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, ip),
		"mkdir -p ~/.ssh && chmod 700 ~/.ssh && cat >> ~/.ssh/authorized_keys")
	pubKey, err := os.ReadFile(identityPubPath)
	if err != nil {
		return fmt.Errorf("failed to read public key from %s: %w", identityPubPath, err)
	}
	cmd.Stdin = bytes.NewBuffer(pubKey)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
