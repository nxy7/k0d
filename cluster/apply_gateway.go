package cluster

import (
	"fmt"
	"k0d/utils"
	"os/exec"
	"strings"
)

func ApplyGateway(gtw string) error {
	return utils.MakeExternalCommandWithStdin(strings.NewReader(GatewayConfig()), "kubectl", "apply", "-f", "/dev/stdin").Run()
}

func AnnotateGateway(ip, ipPool string) error {
	s := utils.MakeSpinner("Annotating Gateway", "Gateway Annotation added")
	s.Start()
	defer s.Stop()
	cmd := exec.Command("kubectl", "apply", "-f", "/dev/stdin")
	cmd.Stdin = strings.NewReader(PoolConfig(ipPool))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	return exec.Command("kubectl", "annotate", "svc", "cilium-gateway-gateway", "-n", "default", fmt.Sprintf("io.cilium/lb-ipam-ips=%s", ip)).Run()
}

func GatewayConfig() string {
	return `apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
metadata:
  name: gateway
  namespace: default
  # issuer is not necesarry for HTTP traffic
  # annotations:
    # "cert-manager.io/cluster-issuer": "ca-issuer"
spec:
  gatewayClassName: cilium
  listeners:
    - protocol: HTTP
      port: 80
      name: chat-app-http
      allowedRoutes:
        namespaces:
          from: All
`
}

func PoolConfig(ipPool string) string {
	return fmt.Sprintf(`apiVersion: "cilium.io/v2alpha1"
kind: CiliumLoadBalancerIPPool
metadata:
  name: "blue-pool"
spec:
  cidrs:
  - cidr: "%s"`, ipPool)
}
