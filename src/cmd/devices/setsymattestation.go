package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"github.com/spf13/cobra"
)

// setSymAttestationCmd represents the device symmetric key attestation create or update command
var setSymAttestationCmd = &cobra.Command{
	Use:   "symmetricKeyAttestation",
	Short: "Create or update a device attestation using symmetric key",
	Long: `Create or update a device attestation using symmetric key`,
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
		primaryKey, err := cmd.Flags().GetString("primaryKey")
		if err != nil {
			return err
		}
		secondaryKey, err := cmd.Flags().GetString("secondaryKey")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// update the device attestation
		p := operations.NewDevicesSetAttestationParams()
		p.DeviceID = id

		p.Body = &models.SymmetricKeyAttestation{
			SymmetricKey: &models.SymmetricKey{
				PrimaryKey:   &primaryKey,
				SecondaryKey: &secondaryKey,
			},
		}
		_, err = c.Operations.DevicesSetAttestation(p)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	setCmd.AddCommand(setSymAttestationCmd)

	setSymAttestationCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setSymAttestationCmd.MarkFlagRequired("app")
	setSymAttestationCmd.Flags().StringP("id", "", "", "unique ID of the device")
	setSymAttestationCmd.MarkFlagRequired("id")
	setSymAttestationCmd.Flags().StringP("primaryKey", "", "", "symmetric primary key")
	setSymAttestationCmd.MarkFlagRequired("primaryKey")
	setSymAttestationCmd.Flags().StringP("secondaryKey", "", "", "symmetric secondary key")
	setSymAttestationCmd.MarkFlagRequired("secondaryKey")
}
