package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start k0d cluster",
	Long:  `Start command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running start command")
	},
}

func init() {
	startCmd.Flags().Int16P("worker-nodes", "w", 1, "")
	startCmd.Flags().Int16P("master-nodes", "m", 1, "")

	startCmd.Flags().BoolP("cilium", "c", true, "")
	startCmd.Flags().BoolP("gateway-api", "g", true, "")
	startCmd.Flags().BoolP("openebs", "o", true, "")
	startCmd.Flags().Bool("certmanager", true, "")
}
