package cluster

import (
	"k0d/utils"
	"os/exec"
)

func InstallOpenEBS() {
	cmd := exec.Command("kubectl", "apply", "-f", "https://openebs.github.io/charts/openebs-operator.yaml")
	err := utils.RunCommandWithSpinner(cmd, "Installing OpenEBS", "OpenEBS Installed\n")
	if err != nil {
		panic(err)
	}
}
