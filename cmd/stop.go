package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop k0d cluster",
	Long:  `Stop command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running stop command")
	},
}

func init() {
	stopCmd.Flags().Bool("cilium", true, "")
}
