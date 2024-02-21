package cluster

import (
	"k0d/utils"
	"log"
	"os"
	"os/exec"
	"sync"
)

func InstallGatewayCrds() {
	s := utils.MakeSpinner("Installing Gateway API Crds v1.0.0", "Gateway CRDs v1.0.0 installed\n")
	s.Start()
	defer s.Stop()
	manifests := []string{
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gatewayclasses.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gateways.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_grpcroutes.yaml",
		"https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_tlsroutes.yaml",
	}
	wg := sync.WaitGroup{}
	for _, manifest := range manifests {
		wg.Add(1)
		go func(manifest string) {
			cmd := exec.Command("kubectl", "apply", "-f", manifest)
			stderr := os.Stderr
			cmd.Stderr = stderr

			err := cmd.Run()
			if err != nil {
				log.Println(manifest)
				panic(err)
			}
			wg.Done()
		}(manifest)
	}
	wg.Wait()
}
