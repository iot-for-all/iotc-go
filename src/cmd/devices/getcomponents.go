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
	"strings"
)

// getComponentsCmd represents the devices get last known value command
var getComponentsCmd = &cobra.Command{
	Use:   "components",
	Short: "Get the list of components in a given device",
	Long:  `Get the list of components in a given device.`,
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

		// Get device components
		p := operations.NewDevicesListComponentsParams()
		p.DeviceID = deviceID
		comps, err := c.Operations.DevicesListComponents(p)
		if err != nil {
			return err
		}

		if len(comps.Payload.Value) == 0 {
			fmt.Printf("No components found in device '%s'\n", deviceID)
		} else {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Type", "Name", "Display Name"})

			numItem := 1
			nextLink := comps.Payload.NextLink
			numItem = addComponentTableRows(t, comps.Payload.Value, numItem)
			// loop through and download all the rows one page at a time
			for {
				if len(nextLink) == 0 {
					break
				}

				spin.Suffix = fmt.Sprintf(" Downloaded %v components, getting more...", numItem-1)
				body, err := util.GetContent(app, nextLink)
				if err != nil {
					return err
				}

				var cc models.Collection
				if err := cc.UnmarshalBinary(body); err != nil {
					return err
				}
				numItem = addComponentTableRows(t, cc.Value, numItem)

				nextLink = cc.NextLink
			}

			spin.Stop()

			// write out the table
			util.RenderTable(t, format, false)
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getComponentsCmd)

	getComponentsCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getComponentsCmd.MarkFlagRequired("app")
	getComponentsCmd.Flags().StringP("id", "", "", "unique device ID")
	getComponentsCmd.MarkFlagRequired("id")
	getComponentsCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

func addComponentTableRows(t table.Writer, components []interface{}, numItem int) int {
	for _, comp := range components {
		cap, ok := comp.(map[string]interface{})
		if ok {
			var compType, compName, compDisplayName  string
			for name, val := range cap {
				switch strings.ToLower(name) {
				case "@type": compType = fmt.Sprintf("%v", val)
				case "name": compName = fmt.Sprintf("%v", val)
				case "displayname": compDisplayName = fmt.Sprintf("%v", val)
				}
			}
			t.AppendRow([]interface{}{numItem, compType, compName, compDisplayName})
			numItem++
		}
	}

	return numItem
}
