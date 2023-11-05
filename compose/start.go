package compose

import (
	"k0d/utils"
	"strings"
)

func Start(config string) error {
	return utils.RunCommandWithSpinner(utils.MakeExternalCommandWithStdin(strings.NewReader(config), "docker", "compose", "-f", "/dev/stdin", "-p", PROJECT_NAME, "up", "-d"), "Creating docker containers...", "Docker containers created\n")
}
