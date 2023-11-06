package utils

import (
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func MakeExternalCommand(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd
}

func MakeExternalCommandWithStdin(in io.Reader, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = in
	return cmd
}

func RunCommandWithSpinner(c *exec.Cmd, spinnerText, finalText string) error {
	// Commands with spinner should have no output right?
	// c.Stdout = nil
	// c.Stderr = nil
	s := MakeSpinner(spinnerText, finalText)
	s.Start()
	defer s.Stop()
	err := c.Run()
	return err
}

func MakeSpinner(spinnerText, finalMessage string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[4], 250*time.Millisecond)
	s.Suffix = "     " + spinnerText
	s.FinalMSG = "✓  " + finalMessage
	return s
}
