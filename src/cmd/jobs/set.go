package jobs

import (
	"bytes"
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// setCmd represents the jobs set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create and execute a new job via its job definition",
	Long: `Create and execute a new job via its job definition.
Job definition needs to the specified in the file using --inputFile.
To run a copy an existing job, follow these steps:
  1. Use the "... jobs get --outputFile" command get the file
  2. Change the 'id' to a unique job id
  3. Remove 'status' field
  4. Pass the file to the "... jobs set --inputFile" command

Sample of input file:
{
  "displayName": "My Job",
  "group": "475cad48-b7ff-4a09-b51e-1a9021385453",
  "data": [
    {
      "type": "PropertyJobData",
      "target": "urn:1dgygpt7t:modelDefinition:02uwtefvdy",
      "path": "componentName.propertyName",
      "value": "updated value"
    }
  ],
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

		// Create a new job
		// TODO: See if the following HTTP PUT with the client API call (i.e. use the JobSet struct)
		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url = url + "/jobs/" + id
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

		fmt.Printf("%s\n", prettyJSON.Bytes())
		return nil
	},
}

func init() {
	jobsCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setCmd.MarkFlagRequired("app")
	setCmd.Flags().StringP("id", "", "", "unique ID for the job")
	setCmd.MarkFlagRequired("id")
	setCmd.Flags().StringP("inputFile", "", "", "file containing the JSON representation of the job")
	setCmd.MarkFlagRequired("inputFile")
}
