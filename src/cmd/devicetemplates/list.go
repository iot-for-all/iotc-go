package devicetemplates

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the deviceTemplates list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the device templates in an application",
	Long:  `List all the device templates in an application`,
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

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading device templates ...")

		// get the list of device templates
		res, err := c.Operations.DeviceTemplatesList(operations.NewDeviceTemplatesListParams())
		if err != nil {
			return err
		}

		if len(res.Payload.Value) == 0 {
			spin.Stop()
			fmt.Printf("No device templates found in '%s' app\n", app)
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Display Name", "Description"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, res.Payload.Value, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := res.Payload.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v device templates, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var dc models.DeviceTemplateCollection
			if err := dc.UnmarshalBinary(body); err != nil {
				return err
			}
			numItem, limitReached, moreRowsExist = addTableRows(t, dc.Value, numItem, top)

			nextLink = dc.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	deviceTemplatesCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func addTableRows(t table.Writer, devices []*models.DeviceTemplate, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range devices {
		t.AppendRow([]interface{}{numItem, item.ID, item.DisplayName, item.Description})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(devices) != i+1
			break
		}
		numItem++
	}
	return numItem, limitReached, moreRowsExist
}

func GetDeviceTemplatesLookupTable(client *client.AzureIoTCentral, app string) (map[string]string, error) {
	deviceTemplates := make(map[string]string)

	res, err := client.Operations.DeviceTemplatesList(operations.NewDeviceTemplatesListParams())
	if err != nil {
		return deviceTemplates, err
	}
	for _, item := range res.Payload.Value {
		deviceTemplates[item.ID] = item.DisplayName
	}

	// see if there are more results beyond first page
	nextLink := res.Payload.NextLink
	for {
		if len(nextLink) == 0 {
			break
		}

		body, err := util.GetContent(app, nextLink)
		if err != nil {
			return deviceTemplates, err
		}

		var dc models.DeviceTemplateCollection
		if err := dc.UnmarshalBinary(body); err != nil {
			return deviceTemplates, err
		}
		for _, item := range dc.Value {
			deviceTemplates[item.ID] = item.DisplayName
		}

		nextLink = dc.NextLink
	}

	return deviceTemplates, nil
}

func GetDeviceTemplate(client *client.AzureIoTCentral, app string, deviceTemplateID string) (*models.DeviceTemplate, error) {

	p := operations.NewDeviceTemplatesGetParams()
	p.DeviceTemplateID = deviceTemplateID
	res, err := client.Operations.DeviceTemplatesGet(p)
	if err != nil {
		return nil, err
	}

	return res.Payload, nil
}
