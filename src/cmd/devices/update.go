package devices

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the devices update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update properties and cloud properties of a device",
	Long: `Update properties and cloud properties of a device.`,
}

func init() {
	devicesCmd.AddCommand(updateCmd)
}
