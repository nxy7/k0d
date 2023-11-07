package cmd

// Maybe this should be --force flag on start
import (
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart k0d cluster",
	Long:  `Restart command`,
	Run: func(cmd *cobra.Command, args []string) {
		stopCmd.Run(cmd, args)
		startCmd.Run(cmd, args)
	},
}

func init() {

	rootCmd.AddCommand(restartCmd)
}
