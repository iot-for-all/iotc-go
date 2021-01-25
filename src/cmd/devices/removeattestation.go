package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/spf13/cobra"
)

// removeAttestationCmd represents the devices remove command
var removeAttestationCmd = &cobra.Command{
	Use:   "removeAttestation",
	Short: "Remove a device attestation by device ID",
	Long: `Remove a device attestation by device ID.`,
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
		spin := util.NewSpinner(" Removing device attestation ...")

		// get the  device by ID
		p := operations.NewDevicesRemoveAttestationParams()
		p.DeviceID = id
		_, err = c.Operations.DevicesRemoveAttestation(p)
		if err != nil {
			return err
		}

		spin.Stop()

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(removeAttestationCmd)

	removeAttestationCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	removeAttestationCmd.MarkFlagRequired("app")
	removeAttestationCmd.Flags().StringP("id", "", "", "unique ID of the device")
	removeAttestationCmd.MarkFlagRequired("id")
}
