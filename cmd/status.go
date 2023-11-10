package cmd

import (
	"fmt"
	"k0d/compose"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check status of k0d cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if compose.IsK0dActive() {
			fmt.Println("✔ There's k0d instance active already")

			// list active cluster utilities like openEBS, GatewayAPI, Cilium
		} else if compose.IsK0dStopped() {
			fmt.Println("There's stopped k0d instance")
		} else {
			fmt.Println("✘ K0d is not running yet")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
