package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"github.com/spf13/cobra"
)

// setCmd represents the devices create or update command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update a device",
	Long: `Create a new device or update an existing one by device ID`,
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
		displayName, err := cmd.Flags().GetString("displayName")
		if err != nil {
			return err
		}
		deviceTemplate, err := cmd.Flags().GetString("deviceTemplate")
		if err != nil {
			return err
		}
		simulated, err := cmd.Flags().GetBool("simulated")
		if err != nil {
			return err
		}
		approved, err := cmd.Flags().GetBool("approved")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// update the device
		p := operations.NewDevicesSetParams()
		p.DeviceID = id
		p.Body = &models.Device{
			DisplayName: displayName,
			InstanceOf: deviceTemplate,
			Simulated: simulated,
			Approved: approved,
		}
		_, err = c.Operations.DevicesSet(p)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setCmd.MarkFlagRequired("app")
	setCmd.Flags().StringP("id", "", "", "unique ID of the device")
	setCmd.MarkFlagRequired("id")
	setCmd.Flags().StringP("displayName", "", "", "display name of the device")
	setCmd.MarkFlagRequired("displayName")
	setCmd.Flags().StringP("deviceTemplate", "", "", "device template (ID) definition for the device")
	setCmd.MarkFlagRequired("deviceTemplate")
	setCmd.Flags().BoolP("simulated", "", false, "whether the device is simulated.")
	setCmd.Flags().BoolP("approved", "", true, "whether the device has been approved to connect to IoT Central")
}
