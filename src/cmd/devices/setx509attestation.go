package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"github.com/spf13/cobra"
)

// setX509AttestationCmd represents the device X509 attestation create or update command
var setX509AttestationCmd = &cobra.Command{
	Use:   "x509Attestation",
	Short: "Create or update a device attestation using X509 cert",
	Long:  `Create or update a device attestation using X509 cert`,
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
		primaryCert, err := cmd.Flags().GetString("primaryCert")
		if err != nil {
			return err
		}
		secondaryCert, err := cmd.Flags().GetString("secondaryCert")
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

		p.Body = &models.X509Attestation{
			X509: &models.X509{
				ClientCertificates: &models.X509Certificates{
					Primary: &models.X509Certificate{
						Certificate: primaryCert,
						Info:        nil,
					},
					Secondary: &models.X509Certificate{
						Certificate: secondaryCert,
						Info:        nil,
					},
				},
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
	setCmd.AddCommand(setX509AttestationCmd)

	setX509AttestationCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setX509AttestationCmd.MarkFlagRequired("app")
	setX509AttestationCmd.Flags().StringP("id", "", "", "unique ID of the device")
	setX509AttestationCmd.MarkFlagRequired("id")
	setX509AttestationCmd.Flags().StringP("primaryCert", "", "", "primary X509 cert")
	setX509AttestationCmd.MarkFlagRequired("primaryCert")
	setX509AttestationCmd.Flags().StringP("secondaryCert", "", "", "secondary X509 cert")
	setX509AttestationCmd.MarkFlagRequired("secondaryCert")
}
