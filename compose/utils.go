package compose

import (
	"os/exec"
	"strings"
)

func IsK0dActive() bool {
	cmd := exec.Command("docker", "compose", "ls")
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	dockerPs := string(b)
	outputByLine := strings.Split(dockerPs, "\n")
	for _, line := range outputByLine[1:] {
		if strings.Contains(line, "k0d") && strings.Contains(line, "running") {
			return true
		}
	}

	return false
}

func IsK0dStopped() bool {
	cmd := exec.Command("docker", "compose", "ls", "-a")
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	dockerPs := string(b)
	outputByLine := strings.Split(dockerPs, "\n")
	for _, line := range outputByLine[1:] {
		if strings.Contains(line, "k0d") && strings.Contains(line, "exited") {
			return true
		}
	}

	return false
}
