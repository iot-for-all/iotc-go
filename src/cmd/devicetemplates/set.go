package devicetemplates

import (
	"bytes"
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// setCmd represents the device templates set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update a device template",
	Long: `Create or update a device template.
Device template definition needs to the specified in the file using --inputFile.

To copy all the existing device templates from another application, you can use
the 'deviceTemplate copy' command.

To copy a single device template from another application, follow these steps:
  1. Use the "... deviceTemplates get -a app1 --outputFile" command to download
     the device template from app1
  2. Remove the 'id' and 'etag' fields from the file
  3. Pass the file to the "... jobs set -a app2 --inputFile" command to upload
     the device template to app2

To update an existing device template in an app, follow these steps:
  1. Use the "... deviceTemplates get -a app1 --outputFile" command to download
     the device template
  2. Remove the 'etag' field from the file
  3. Pass the file to the "... jobs set -a app1 --inputFile" command to upload
     the device template

To create a new device template in an app, pass in the device template JSON file
similar to the one you get from 'get' command
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
		spin := util.NewSpinner(" Updating device template ...")

		// Create a new device template
		// TODO: See if the following HTTP PUT with the client API call (i.e. use the DeviceTemplateSet struct)
		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url = url + "/deviceTemplates/" + id
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
	deviceTemplatesCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setCmd.MarkFlagRequired("app")
	setCmd.Flags().StringP("id", "", "", "unique ID for the device template")
	setCmd.MarkFlagRequired("id")
	setCmd.Flags().StringP("inputFile", "", "", "file containing the JSON representation of the job")
	setCmd.MarkFlagRequired("inputFile")
}
