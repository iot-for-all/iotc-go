package devices

import (
	"com.azure.iot/iotcentral/iotcgo/models"
	"fmt"
	"os"

	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// getModulesCmd represents the devices get list of modules command
var getModulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Get the list of modules in a given edge device",
	Long:  `Get the list of modules in a given edge device.`,
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
		spin := util.NewSpinner(" Downloading device ...")

		// get the device by ID
		p := operations.NewDevicesListModulesParams()
		p.DeviceID = id
		res, err := c.Operations.DevicesListModules(p)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Type", "Name", "Display Name"})
		numItem := 1
		numItem = addModelTableRows(t, res.Payload.Value, numItem)
		nextLink := res.Payload.NextLink

		for {
			if len(nextLink) == 0 {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v modules, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var mc models.Collection
			if err := mc.UnmarshalBinary(body); err != nil {
				return err
			}
			numItem  = addModelTableRows(t, mc.Value, numItem)

			nextLink = mc.NextLink
		}

		spin.Stop()

		util.RenderTable(t, format, false)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getModulesCmd)

	getModulesCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getModulesCmd.MarkFlagRequired("app")
	getModulesCmd.Flags().StringP("id", "", "", "unique device ID")
	getModulesCmd.MarkFlagRequired("id")
	getModulesCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

func addModelTableRows(t table.Writer, modules []interface{}, numItem int) int {
	for _, mod := range modules {
		module, ok := mod.(map[string]interface{})
		if ok {
			types, ok := module["@type"].([]interface{})
			typeStr := ""
			if ok {
				for _, typ := range types {
					if len(typeStr) > 0 {
						typeStr += ", "
					}
					typeStr += typ.(string)
				}
			}
			t.AppendRow([]interface{}{numItem, typeStr, module["name"], module["displayName"]})
		}

		numItem++
	}

	return numItem
}

