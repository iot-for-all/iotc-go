package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicetemplates"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the devices list command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details about an existing device by device ID",
	Long: `Get details about an existing device by device ID.`,
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
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading device ...")

		// get the  device by ID
		p := operations.NewDevicesGetParams()
		p.DeviceID = id
		res, err := c.Operations.DevicesGet(p)
		if err != nil {
			return err
		}

		spin.Suffix = " Downloading device template ..."
		template, err := devicetemplates.GetDeviceTemplate(c, app, res.Payload.InstanceOf)
		if err != nil {
			return err
		}

		spin.Stop()
		printDevice(res.Payload, template, app, format)

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCmd.MarkFlagRequired("app")
	getCmd.Flags().StringP("id", "", "", "unique ID of the device")
	getCmd.MarkFlagRequired("id")
	getCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}


// printDevice prints the device
func printDevice(device *models.Device, template *models.DeviceTemplate, app string, format string) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", device.ID})
	t.AppendRow([]interface{}{"Display Name", device.DisplayName})
	t.AppendRow([]interface{}{"Device Template ID", device.InstanceOf})
	t.AppendRow([]interface{}{"Device Template Name", template.DisplayName})
	t.AppendRow([]interface{}{"Provisioned", *device.Provisioned})
	t.AppendRow([]interface{}{"Approved", device.Approved})
	t.AppendRow([]interface{}{"Simulated", device.Simulated})

	util.RenderTable(t, format, false)
}