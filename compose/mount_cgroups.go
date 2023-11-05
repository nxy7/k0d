package compose

import (
	"fmt"
	"os/exec"
)

func MountCgroups() {
	err := exec.Command("sudo", "mkdir", "-p", "/run/cilium/cgroupv2").Run()
	if err != nil {
		panic(err)
	}
	err = exec.Command("sudo", "mount", "-t", "cgroup2", "none", "/run/cilium/cgroupv2").Run()
	if err != nil {
		fmt.Println("Cgroups already mounted (?)")
	}
	err = exec.Command("sudo", "mount", "--make-shared", "/run/cilium/cgroupv2").Run()
	if err != nil {
		panic(err)
	}
	err = exec.Command("echo", "Done mounting cgroupv2").Run()
	if err != nil {
		panic(err)
	}
}
