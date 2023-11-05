package cluster

import (
	"k0d/utils"
	"os/exec"
	"strings"
)

func InstallCertManager() {
	s := utils.MakeSpinner("Installing cert manager", "Cert manager installed\n")
	s.Start()
	defer s.Stop()
	err := exec.Command("helm", "repo", "add", "jetstack", "https://charts.jetstack.io").Run()
	if err != nil {
		panic(err)
	}

	err = exec.Command("helm", "repo", "update").Run()
	if err != nil {
		panic(err)
	}

	err = exec.Command("helm", "install", "cert-manager",
		"--version", "v1.13.0",
		"--namespace", "cert-manager",
		"--create-namespace",
		"--set", "installCRDs=true",
		"--set", `extraArgs={--feature-gates=ExperimentalGatewayAPISupport=true}`,
		"jetstack/cert-manager").Run()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("kubectl", "apply", "-f", "/dev/stdin")
	cmd.Stdin = strings.NewReader(utils.MakeCaIssuerConfig())
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
