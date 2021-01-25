package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"github.com/spf13/cobra"
)

// setTpmAttestationCmd represents the device TPM attestation create or update command
var setTpmAttestationCmd = &cobra.Command{
	Use:   "tpmAttestation",
	Short: "Create or update a device attestation using TPM",
	Long: `Create or update a device attestation using TPM`,
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
		endorsementKey, err := cmd.Flags().GetString("endorsementKey")
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

		p.Body = &models.TpmAttestation{
			Tpm: &models.Tpm{
				EndorsementKey: &endorsementKey,
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
	setCmd.AddCommand(setTpmAttestationCmd)

	setTpmAttestationCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setTpmAttestationCmd.MarkFlagRequired("app")
	setTpmAttestationCmd.Flags().StringP("id", "", "", "unique ID of the device")
	setTpmAttestationCmd.MarkFlagRequired("id")
	setTpmAttestationCmd.Flags().StringP("endorsementKey", "", "", "endorsement key of the TPM")
	setTpmAttestationCmd.MarkFlagRequired("endorsementKey")
}
