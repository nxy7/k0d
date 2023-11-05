package compose

import "k0d/utils"

func Stop() error {
	return utils.MakeExternalCommand("docker", "compose", "-p", PROJECT_NAME, "down").Run()
}
