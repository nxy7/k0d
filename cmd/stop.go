package cmd

import (
	"k0d/compose"
	"k0d/utils"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop k0d cluster",
	Long:  `Stop command`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.RunCommandWithSpinner(utils.MakeExternalCommand("docker", "compose", "-p", compose.PROJECT_NAME, "stop"), "Shutting k0d down..", "Done!\n")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	stopCmd.Flags().Bool("cilium", true, "")

	rootCmd.AddCommand(stopCmd)
}
