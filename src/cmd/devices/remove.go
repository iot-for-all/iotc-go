package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/spf13/cobra"
)

// removeCmd represents the devices remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an existing device by device ID",
	Long: `Remove an existing device by device ID.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		app, err := cmd.Flags().GetString("app")
		if err != nil {
			return err
		}
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Removing device ...")

		// get the  device by ID
		p := operations.NewDevicesRemoveParams()
		p.DeviceID = id
		_, err = c.Operations.DevicesRemove(p)
		if err != nil {
			return err
		}

		spin.Stop()

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	removeCmd.MarkFlagRequired("app")
	removeCmd.Flags().StringP("id", "", "", "unique ID of the device")
	removeCmd.MarkFlagRequired("id")
}
