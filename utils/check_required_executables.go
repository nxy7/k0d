package utils

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func CheckRequiredExecutables(c *cobra.Command) {
	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("K0d needs 'docker' executable in path")
		panic(err)
	}
	_, err = exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("K0d needs 'kubectl' executable in path")
		panic(err)
	}
	_, err = exec.LookPath("helm")
	if err != nil {
		fmt.Println("K0d needs 'helm' CLI utility in path")
		panic(err)
	}

	ciliumRequired, err := c.Flags().GetBool("cilium")
	if err != nil {
		panic(err)
	}
	if ciliumRequired {
		_, err := exec.LookPath("cilium")
		if err != nil {
			fmt.Println("To use cilium CNI with k0d you need to have cilium CLI utility installed.")
			panic(err)
		}
	}
}
