package cde

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// cdeCmd represents the CDE command
var cdeCmd = &cobra.Command{
	Use:   "cde",
	Short: "Manage data exports within your IoT Central application",
	Long: `Continuous Data Export lets you export data from Central to different data sinks.
Manage data exports within your IoT Central application.`,
}

func init() {
	cmd.RootCmd.AddCommand(cdeCmd)
}