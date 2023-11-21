package cmd

import (
	"k0d/compose"
	"k0d/utils"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "kill k0d cluster",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.RunCommandWithSpinner(utils.MakeExternalCommand("docker", "compose", "-p", compose.PROJECT_NAME, "down"), "Stopping k0d down..", "Done!\n")
		if err != nil {
			panic(err)
		}
	},
}

func init() {

	rootCmd.AddCommand(killCmd)
}
