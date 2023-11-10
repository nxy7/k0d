package cluster

import (
	"bytes"
	"fmt"
	"k0d/utils"
	"os/exec"
)

func InstallCillium() {
	p, err := utils.SaveToFile(CiliumValues(), "ciliumValues.yaml")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("cilium", "install", "--values", p, "--version", "1.14.3", "--wait")
	var serr bytes.Buffer
	cmd.Stderr = &serr

	err = utils.RunCommandWithSpinner(cmd, "Installing Cillium", "Cilium Installed\n")
	if err != nil {
		fmt.Println(serr.String())
		panic(err)
	}
}

func CiliumValues() string {
	return `kubeProxyReplacement: true
k8sServiceHost: 127.0.0.1
k8sServicePort: 6443
operator:
  replicas: 1
envoy:
  enabled: true
debug:
  enable: true
  verbose: envoy
l2announcements:
  enabled: true
bgpControlPlane:
  enabled: true
devices: "eth0"
externalIPs:
  enabled: true
# nodePort:
#  enabled: true
gatewayAPI:
  enabled: true
ipam:
  mode: kubernetes
# bpf:
#  masquerade: true
#  hostLegacyRouting: false
ipv4:
  enabled: true
# ipv6:
#  enabled: false
hubble:
  enabled: true
  ui:
    enabled: true
  relay:
    enabled: true`
}
