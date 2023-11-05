package cluster

import (
	"k0d/utils"
	"os/exec"
	"sync"
)

func InstallGatewayCrds() {
	s := utils.MakeSpinner("Installing Gateway API Crds", "Gateway CRDs installed\n")
	s.Start()
	defer s.Stop()
	manifests := []string{
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v0.8.1/config/crd/standard/gateway.networking.k8s.io_gatewayclasses.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v0.8.1/config/crd/standard/gateway.networking.k8s.io_gateways.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v0.8.1/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v0.8.1/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v0.8.1/config/crd/experimental/gateway.networking.k8s.io_tlsroutes.yaml",
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
