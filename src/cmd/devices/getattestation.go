package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getAttestationCmd represents the devices get attestation command
var getAttestationCmd = &cobra.Command{
	Use:   "attestation",
	Short: "Get the attestation of a given device",
	Long:  `Get the attestation of a given device.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		app, err := cmd.Flags().GetString("app")
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString("format")
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
		spin := util.NewSpinner(" Downloading device attestation ...")

		// get the  device by ID
		p := operations.NewDevicesGetAttestationParams()
		p.DeviceID = id
		res, err := c.Operations.DevicesGetAttestation(p)
		if err != nil {
			return err
		}

		spin.Stop()
		printAttestation(res.Payload, format)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAttestationCmd)

	getAttestationCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getAttestationCmd.MarkFlagRequired("app")
	getAttestationCmd.Flags().StringP("id", "", "", "unique device ID")
	getAttestationCmd.MarkFlagRequired("id")
	getAttestationCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

func printAttestation(attestation models.Attestation, format string) error {
	//TODO print different types of attestations
	fmt.Printf("Attestation: %v", attestation)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Type"})

	t.AppendRow([]interface{}{attestation.Type()})

	util.RenderTable(t, format, false)
	return nil
}
