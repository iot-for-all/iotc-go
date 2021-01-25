package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicetemplates"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCommandsCmd represents the devices get list of commands command
var getCommandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Get the list of commands in a given device",
	Long:  `Get the list of commands in a given device.`,
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

		commands := util.GetComponentCommands(template)

		spin.Stop()
		printCommands(commands, format)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getCommandsCmd)

	getCommandsCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCommandsCmd.MarkFlagRequired("app")
	getCommandsCmd.Flags().StringP("id", "", "", "unique device ID")
	getCommandsCmd.MarkFlagRequired("id")
	getCommandsCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

func printCommands(commands []util.ComponentCommands, format string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Component", "Command"})

	numRow := 1
	for _, cc := range commands {
		for _, command := range cc.Commands {
			t.AppendRow([]interface{}{numRow, cc.Component, command})
			numRow++
		}
	}

	util.RenderTable(t, format, false)
}