package cde

import (
	"com.azure.iot/iotcentral/iotcgo/config"
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the cde list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of continuous data exports in an application.",
	Long: `Get the list of continuous data exports in an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		// read the command line parameters
		app, err := cmd.Flags().GetString("app")
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// get the list of CDE exports
		res, err := c.Operations.ContinuousDataExportsList(operations.NewContinuousDataExportsListParams())
		if err != nil {
			return err
		}

		printTable(res.Payload.Value, app, format)
		 */

		fmt.Println("CDE V2 is not yet supported in IoT Central API.")
		return nil
	},
}

func init() {
	cdeCmd.AddCommand(listCmd)

	listCmd.Flags().StringP( "app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

/*
// printTable prints the CDE list as a table
func printTable(cde []*models.ContinuousDataExport, app string, format string) {
	if len(cde) == 0 {
		fmt.Printf("No CDE exports found in '" + app + "' app\n")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "Display Name", "Status", "Source", "Enabled"})

	for i, item := range cde {
		t.AppendRow([]interface{}{i + 1, item.ID, item.DisplayName, item.Status, item.Sources, item.Enabled})
	}
	util.RenderTable(t, format)
}
*/
