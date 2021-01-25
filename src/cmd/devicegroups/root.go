package devicegroups

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// deviceGroupsCmd represents the deviceGroups command
var deviceGroupsCmd = &cobra.Command{
	Use:   "deviceGroups",
	Short: "Operate against IoT Central device groups",
	Long: `Device Groups lets you group a set of devices based on a query.`,
}

func init() {
	cmd.RootCmd.AddCommand(deviceGroupsCmd)
}