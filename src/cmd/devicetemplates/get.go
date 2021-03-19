package devicetemplates

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// getCmd represents the device templates get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a device template by ID",
	Long: `Get a device template by ID.
If the '--outputFile' is specified, JSON representation of the given 
device template is written to the file. This JSON file can be easily
edited as a template to be used in the 'set' command.`,
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
		outputFile, err := cmd.Flags().GetString("outputFile")
		if err != nil {
			return err
		}
		getMerged, err := cmd.Flags().GetBool("getMerged")
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading device template ...")

		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/deviceTemplates/" + id
		if getMerged {
			url += "/merged"
		}
		body, err := util.GetIndentedJSONContent(app, url) // get the
		if err != nil {
			return err
		}

		// parse JSON into the device template struct
		var dt models.DeviceTemplate
		if err := dt.UnmarshalBinary(body); err != nil {
			return err
		}

		spin.Stop()

		// if the output file is given, write the contents into a file
		if len(outputFile) > 0 {
			err = ioutil.WriteFile(outputFile, body, 0755)
			if err != nil {
				return err
			}

			fmt.Printf("'%s' Device template is written into the file: %s\n", dt.DisplayName, outputFile)
		} else {
			fmt.Printf("%s\n", body)
		}

		return nil
	},
}

func init() {
	deviceTemplatesCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCmd.MarkFlagRequired("app")
	getCmd.Flags().StringP("id", "", "", "unique ID for the device template")
	getCmd.MarkFlagRequired("id")
	getCmd.Flags().StringP("outputFile", "", "", "dump the device template JSON to the given file")
	getCmd.Flags().BoolP("getMerged", "", false, "get the solution model (overrides, initial values) merged into capability model")
}
