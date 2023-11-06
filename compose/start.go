package compose

import (
	"k0d/utils"
	"os/exec"
	"strings"
)

func Start(config string) error {
	cmd := exec.Command("docker", "compose", "-f", "/dev/stdin", "-p", PROJECT_NAME, "up", "-d")
	cmd.Stdin = strings.NewReader(config)
	return utils.RunCommandWithSpinner(cmd, "Creating docker containers...", "Docker containers created\n")
}
