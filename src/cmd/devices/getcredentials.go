package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCredentialsCmd represents the devices get credentials command
var getCredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Get the credentials of a given device",
	Long:  `Get the credentials of a given device.`,
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
		deviceID, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Getting device components ...")

		// Get device credentials
		p := operations.NewDevicesGetCredentialsParams()
		p.DeviceID = deviceID
		credentials, err := c.Operations.DevicesGetCredentials(p)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Scope ID", "Primary Key", "Secondary Key"})
		t.AppendRow([]interface{}{*credentials.Payload.IDScope,
			*credentials.Payload.SymmetricKey.PrimaryKey,
			*credentials.Payload.SymmetricKey.SecondaryKey})
		spin.Stop()

		// write out the table
		util.RenderTable(t, format, false)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getCredentialsCmd)

	getCredentialsCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCredentialsCmd.MarkFlagRequired("app")
	getCredentialsCmd.Flags().StringP("id", "", "", "unique device ID")
	getCredentialsCmd.MarkFlagRequired("id")
	getCredentialsCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

