package cmd

import (
	"fmt"
	"k0d/cluster"
	"k0d/compose"
	"k0d/utils"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start k0d cluster",
	Long:  `Start command`,
	Run: func(cmd *cobra.Command, args []string) {
		if compose.IsK0dActive() {
			fmt.Println("There's k0d instance active already")
			return
		}
		if compose.IsK0dStopped() {
			fmt.Println("Restarting stopped k0d instance")
			compose.Start(compose.MakeComposeFile(cmd))
			return
		}
		utils.CheckRequiredExecutables(cmd)

		composeConfig := compose.MakeComposeFile(cmd)
		err := compose.Start(composeConfig)
		if err != nil {
			panic(err)
		}

		utils.WaitForCluster()
		compose.MountCgroups()
		cluster.InstallGatewayCrds()

		cluster.InstallCillium()

		cluster.InstallOpenEBS()
		cluster.InstallCertManager()

		cluster.ApplyGateway(cluster.GatewayConfig())
		cluster.AnnotateGateway("172.17.0.2", "172.17.0.2/24")

	},
}

func init() {
	startCmd.Flags().Int16P("worker-nodes", "w", 1, "")
	startCmd.Flags().Int16P("master-nodes", "m", 1, "")

	startCmd.Flags().BoolP("cilium", "c", true, "")
	startCmd.Flags().BoolP("gateway-api", "g", true, "")
	startCmd.Flags().BoolP("openebs", "o", true, "")
	startCmd.Flags().Bool("certmanager", true, "")

	rootCmd.AddCommand(startCmd)
}
