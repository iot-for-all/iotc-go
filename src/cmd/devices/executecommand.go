package devices

import (
	"bytes"
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// executeCommandCmd represents the devices execute Command command
var executeCommandCmd = &cobra.Command{
	Use:   "executeCommand",
	Short: "Execute command on the device",
	Long: `Execute command on the device.
To execute a command belonging to a component, pass the component name
argument. To execute command belonging to the default component, you
can skip the component argument. To execute commands in a module pass
in the module argument, otherwise skip it.

You can use the 'devices get commands' to get the list of commands.

E.g.: Properties JSON file
{
  "request": {
    "tempVal": 30,
    "param2": "value2"
  },
}
`,
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
		module, err := cmd.Flags().GetString("module")
		if err != nil {
			return err
		}
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			return err
		}
		inputFile, err := cmd.Flags().GetString("inputFile")
		if err != nil {
			return err
		}
		component, err := cmd.Flags().GetString("component")
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Executing command ...")

		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/devices/" + id
		if module != "" {
			url += "/modules/" + module
		}
		if component != "" {
			url += "/components/" + component
		}
		url += "/commands/" + command

		// execute command
		body, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return err
		}
		response, err := util.PostContent(app, url, body)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, response, "", "  ")
		if err != nil {
			return err
		}

		spin.Stop()
		fmt.Printf("%s\n", prettyJSON.Bytes())

		return nil
	},
}

func init() {
	devicesCmd.AddCommand(executeCommandCmd)

	executeCommandCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	executeCommandCmd.MarkFlagRequired("app")
	executeCommandCmd.Flags().StringP("id", "", "", "unique ID of the device")
	executeCommandCmd.MarkFlagRequired("id")
	executeCommandCmd.Flags().StringP("command", "", "", "name of the command")
	executeCommandCmd.MarkFlagRequired("command")
	executeCommandCmd.Flags().StringP("inputFile", "", "", "JSON file containing the command parameters")
	executeCommandCmd.MarkFlagRequired("inputFile")
	executeCommandCmd.Flags().StringP("component", "", "", "the name of the component containing the command")
	executeCommandCmd.Flags().StringP("module", "", "", "the name of the module containing the command")
}
