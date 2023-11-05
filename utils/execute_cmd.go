package utils

import (
	"io"
	"os"
	"os/exec"
)

func RunExternalCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func RunExternalCommandWithStdin(in io.Reader, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = in
	return cmd.Run()
}
