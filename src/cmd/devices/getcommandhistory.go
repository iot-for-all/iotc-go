package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCommandHistoryCmd represents the devices get list of command history command
var getCommandHistoryCmd = &cobra.Command{
	Use:   "commandHistory",
	Short: "Get the list of command history in a given device",
	Long:  `Get the list of command history in a given device.`,
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
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			return err
		}
		component, err := cmd.Flags().GetString("component")
		if err != nil {
			return err
		}
		module, err := cmd.Flags().GetString("module")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading command history ...")

		// get the  device by ID
		var history []*models.DeviceCommand
		if module != "" {
			if command != "" {
				p := operations.NewDevicesGetModuleComponentCommandHistoryParams()
				p.DeviceID = id
				p.CommandName = command
				p.ComponentName = component
				p.ModuleName = module
				res, err := c.Operations.DevicesGetModuleComponentCommandHistory(p)
				if err != nil {
					return err
				}
				history = res.Payload.Value
			} else {
				p := operations.NewDevicesGetModuleCommandHistoryParams()
				p.DeviceID = id
				p.CommandName = command
				p.ModuleName = module
				res, err := c.Operations.DevicesGetModuleCommandHistory(p)
				if err != nil {
					return err
				}
				history = res.Payload.Value
			}
		} else {
			if command != "" {
				p := operations.NewDevicesGetComponentCommandHistoryParams()
				p.DeviceID = id
				p.CommandName = command
				p.ComponentName = component
				res, err := c.Operations.DevicesGetComponentCommandHistory(p)
				if err != nil {
					return err
				}
				history = res.Payload.Value
			} else {
				p := operations.NewDevicesGetCommandHistoryParams()
				p.DeviceID = id
				p.CommandName = command
				res, err := c.Operations.DevicesGetCommandHistory(p)
				if err != nil {
					return err
				}
				history = res.Payload.Value
			}
		}

		// TODO: get more rows if needed
		// Currently it is returning only one row

		spin.Stop()
		if err := printCommandHistory(history, format); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getCommandHistoryCmd)

	getCommandHistoryCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCommandHistoryCmd.MarkFlagRequired("app")
	getCommandHistoryCmd.Flags().StringP("id", "", "", "unique device ID")
	getCommandHistoryCmd.MarkFlagRequired("id")
	getCommandHistoryCmd.Flags().StringP("command", "", "", "name of the command")
	getCommandHistoryCmd.MarkFlagRequired("command")
	getCommandHistoryCmd.Flags().StringP("component", "", "", "the name of the component containing the command")
	getCommandHistoryCmd.Flags().StringP("module", "", "", "the name of the module containing the command")
	getCommandHistoryCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

func printCommandHistory(commands []*models.DeviceCommand, format string) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Request", "Response", "Response Code"})

	numRow := 1
	for _, cc := range commands {
		jsonBytes, err := json.Marshal(cc.Request)
		if err != nil {
			return err
		}
		requestJson := string(jsonBytes)
		var responseJson = ""
		if cc.Response != nil {
			jsonBytes, err = json.Marshal(cc.Response)
			if err != nil {
				return err
			}
			responseJson = string(jsonBytes)
		}
		t.AppendRow([]interface{}{numRow, requestJson, responseJson, cc.ResponseCode})
		numRow++
	}

	util.RenderTable(t, format, false)
	return nil
}
