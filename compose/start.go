package compose

import (
	"k0d/utils"
	"strings"
)

func Start(config string) error {
	return utils.RunExternalCommandWithStdin(strings.NewReader(config), "docker", "compose", "-f", "/dev/stdin", "-p", PROJECT_NAME, "up", "-d")
}
