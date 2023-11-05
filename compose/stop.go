package compose

import "k0d/utils"

func Stop() error {
	return utils.RunExternalCommand("docker", "compose", "-p", PROJECT_NAME, "down")
}
