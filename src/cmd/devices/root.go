package devices

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Get information about and manage devices and IoT Edge modules in your IoT Central application.",
	Long: `Get information about and manage devices and IoT Edge modules in your IoT Central application.`,
}

func init() {
	cmd.RootCmd.AddCommand(devicesCmd)
}