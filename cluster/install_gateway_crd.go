package cluster

import (
	"k0d/utils"
	"os/exec"
	"sync"
)

func InstallGatewayCrds() {
	s := utils.MakeSpinner("Installing Gateway API Crds v1.0.0", "Gateway CRDs v1.0.0 installed\n")
	s.Start()
	defer s.Stop()
	manifests := []string{
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gatewayclasses.yaml",
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gateways.yaml",
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml",
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml",
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_grpcroutes.yaml",
		"kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_tlsroutes.yaml",
	}
	wg := sync.WaitGroup{}
	for _, manifest := range manifests {
		wg.Add(1)
		go func(manifest string) {
			err := exec.Command("kubectl", "apply", "-f", manifest).Run()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(manifest)
	}
	wg.Wait()
}
