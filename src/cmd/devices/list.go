package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicetemplates"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// listCmd represents the devices list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the devices in an application",
	Long:  `List all the devices in an application`,
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
		top, err := cmd.Flags().GetInt("top")
		if err != nil {
			return err
		}
		deviceTemplateID, err := cmd.Flags().GetString("deviceTemplate")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading devices ...")

		// get the list of devices
		var devices []*models.Device
		var nextLink string
		if strings.ToLower(deviceTemplateID) == "all" {
			res, err := c.Operations.DevicesList(operations.NewDevicesListParams())
			if err != nil {
				return err
			}

			devices = res.Payload.Value
			nextLink = res.Payload.NextLink
		} else {
			p := operations.NewDeviceTemplatesListDevicesParams()
			p.DeviceTemplateID = deviceTemplateID
			res, err := c.Operations.DeviceTemplatesListDevices(p)
			if err != nil {
				return err
			}

			devices = res.Payload.Value
			nextLink = res.Payload.NextLink
		}

		if len(devices) == 0 {
			spin.Stop()
			fmt.Printf("No devices found in '%s' app\n", app)
			return nil
		}

		// get all device templates look up table so that we can print the template names
		spin.Suffix = " Getting device templates"
		deviceTemplates, err := devicetemplates.GetDeviceTemplatesLookupTable(c, app)
		if err != nil {
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Display Name", "Device Template", "Provisioned", "Approved", "Simulated"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, devices, deviceTemplates, numItem, top)

		// loop through and download all the rows one page at a time
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v devices, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var dc models.DeviceCollection
			if err := dc.UnmarshalBinary(body); err != nil {
				return err
			}
			numItem, limitReached, moreRowsExist = addTableRows(t, dc.Value, deviceTemplates, numItem, top)

			nextLink = dc.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
	listCmd.Flags().StringP("deviceTemplate", "", "all", "list devices of the specified device template ID")
}

func addTableRows(t table.Writer, devices []*models.Device, deviceTemplates map[string]string, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range devices {
		t.AppendRow([]interface{}{numItem, item.ID, item.DisplayName, deviceTemplates[item.InstanceOf], *item.Provisioned, item.Approved, item.Simulated})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(devices) != i+1
			break
		}
		numItem++
	}
	return numItem, limitReached, moreRowsExist
}