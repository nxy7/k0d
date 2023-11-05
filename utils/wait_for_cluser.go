package utils

import (
	"os/exec"
	"time"
)

func WaitForCluster() {
	s := MakeSpinner("Waiting till cluster becomes available..", "Cluster ready\n")
	s.Start()
	defer s.Stop()
	for {
		err := CopyKubeconfig()
		if err == nil {
			err = exec.Command("kubectl", "get", "svc").Run()
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second)
	}
}
