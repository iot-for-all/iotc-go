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

// updateCloudPropertiesCmd represents the devices update cloud properties command
var updateCloudPropertiesCmd = &cobra.Command{
	Use:   "cloudProperties",
	Short: "Update device cloud properties",
	Long: `Update device cloud properties.
You can use the 'devices get lkv' command to get current values.

E.g.: Cloud properties JSON file
{
  "CustomerName": "Acme corp",
  "City": "Redmond",
  "State": "WA"
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

		// start the spinner
		spin := util.NewSpinner(" Updating properties ...")

		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/devices/" + id + "/cloudProperties"

		// update cloud properties
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
	updateCmd.AddCommand(updateCloudPropertiesCmd)

	updateCloudPropertiesCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	updateCloudPropertiesCmd.MarkFlagRequired("app")
	updateCloudPropertiesCmd.Flags().StringP("id", "", "", "unique ID of the device")
	updateCloudPropertiesCmd.MarkFlagRequired("id")
	updateCloudPropertiesCmd.Flags().StringP("inputFile", "", "", "JSON file containing the cloud properties")
	updateCloudPropertiesCmd.MarkFlagRequired("inputFile")
}
