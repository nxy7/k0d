package cmd

// Maybe this should be --force flag on start
import (
	"fmt"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart k0d cluster",
	Long:  `Restart command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running start command")
	},
}

func init() {
	restartCmd.Flags().Int16P("worker-nodes", "w", 1, "")
	restartCmd.Flags().Int16P("master-nodes", "m", 1, "")

	restartCmd.Flags().BoolP("cilium", "c", true, "")
	restartCmd.Flags().BoolP("gateway-api", "g", true, "")
	restartCmd.Flags().BoolP("openebs", "o", true, "")
	restartCmd.Flags().Bool("certmanager", true, "")
}
