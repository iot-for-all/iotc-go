package apps

import (
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the apps get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the app details from the given subscription",
	Long: `Get the app details from the given subscription.
You need to specify the application resource name to get. You can get it
using the 'apps list' command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		subscription, err := cmd.Flags().GetString("subscription")
		if err != nil {
			return err
		}
		resourceGroup, err := cmd.Flags().GetString("resourceGroup")
		if err != nil {
			return err
		}
		resourceName, err := cmd.Flags().GetString("resourceName")
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner("  Connecting to IoT Central...")

		// create a new client
		client := iotcentral.NewAppsClient(subscription)

		// create an authorizer from az cli
		authorizer, err := auth.NewAuthorizerFromCLI()
		if err != nil {
			return err
		}
		client.Authorizer = authorizer

		// get the app details
		spin.Suffix = " Downloading IoT Central app details..."
		app, err := client.Get(context.Background(), resourceGroup, resourceName)
		if err != nil {
			return err
		}

		// write out the table
		spin.Stop()
		printApplication(app, format)

		return nil
	},
}

func init() {
	appsCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("subscription", "s", "", "Azure subscription ID")
	getCmd.MarkFlagRequired("subscription")
	getCmd.Flags().StringP("resourceGroup", "", "", "the resource group ID of the application")
	getCmd.MarkFlagRequired("resourceGroup")
	getCmd.Flags().StringP("resourceName", "a", "", "the name of the application resource")
	getCmd.MarkFlagRequired("resourceName")
	getCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

// printApplication prints the app details as a table
func printApplication(app iotcentral.App, format string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", *app.ApplicationID})
	t.AppendRow([]interface{}{"App Name", *app.DisplayName})
	t.AppendRow([]interface{}{"Subdomain", *app.Subdomain})
	t.AppendRow([]interface{}{"Resource Name", *app.Name})
	t.AppendRow([]interface{}{"Location", *app.Location})
	t.AppendRow([]interface{}{"SKU", app.Sku.Name})
	t.AppendRow([]interface{}{"Template", *app.Template})
	t.AppendRow([]interface{}{"Type", *app.Type})
	var tagStr string
	for name, val := range app.Tags {
		tagStr += fmt.Sprintf("{%s: %s} ", name, *val)
	}
	t.AppendRow([]interface{}{"Tags", tagStr})

	util.RenderTable(t, format, false)
}
