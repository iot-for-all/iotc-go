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

// updatePropertiesCmd represents the devices update properties command
var updatePropertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "Update device twin properties",
	Long: `Update device twin properties.
To update properties belonging to a component, pass the component name
argument. To update properties belonging to the default component, you
can skip the component argument.

You can use the 'devices get lkv' command to get current values.

E.g.: Properties JSON file
{
  "fanSpeed": 35,
  "voltage": 5,
  "current": 2,
  "irSwitch": true
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
		inputFile, err := cmd.Flags().GetString("inputFile")
		if err != nil {
			return err
		}
		component, err := cmd.Flags().GetString("component")
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Updating properties ...")

		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/devices/" + id
		if component != "" {
			url += "/components/" + component
		}
		url += "/properties"

		// update properties
		body, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return err
		}
		response, err := util.PutContent(app, url, body)
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
	updateCmd.AddCommand(updatePropertiesCmd)

	updatePropertiesCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	updatePropertiesCmd.MarkFlagRequired("app")
	updatePropertiesCmd.Flags().StringP("id", "", "", "unique ID of the device")
	updatePropertiesCmd.MarkFlagRequired("id")
	updatePropertiesCmd.Flags().StringP("inputFile", "", "", "JSON file containing the properties")
	updatePropertiesCmd.MarkFlagRequired("inputFile")
	updatePropertiesCmd.Flags().StringP("component", "", "", "the name of the component containing the command")
}
