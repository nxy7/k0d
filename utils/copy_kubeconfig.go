package utils

import (
	"os"
	"os/exec"
)

// Copies kubeconfig from cluster into working machine ~/.kube/config
func CopyKubeconfig() error {
	// docker exec k0s cat /var/lib/k0s/pki/admin.conf
	cmd := exec.Command("docker", "exec", "k0s", "cat", "/var/lib/k0s/pki/admin.conf")
	config, err := cmd.Output()
	if err != nil {
		return err
	}

	// save ~/.kube/config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = os.WriteFile(homeDir+"/.kube/config", config, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
