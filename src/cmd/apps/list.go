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

// listCmd represents the apps list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all the apps in the given subscription",
	Long: `Get the list of all the apps in the given subscription.`,
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

		// list all apps in the subscription
		spin.Suffix = " Downloading IoT Central apps list..."
		var apps iotcentral.AppListResultPage
		if len(resourceGroup) != 0 {
			apps, err = client.ListByResourceGroup(context.Background(), resourceGroup)
		} else {
			apps, err = client.ListBySubscription(context.Background())
		}
		if err != nil {
			return err
		}

		// write out the table
		spin.Stop()
		printTable(apps.Values(), format)

		return nil
	},
}

func init() {
	appsCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("subscription", "s", "", "Azure subscription ID")
	listCmd.MarkFlagRequired("subscription")
	listCmd.Flags().StringP("resourceGroup", "", "", "the resource group in which the apps are created")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

// printTable prints the apps list as a table
func printTable(apps []iotcentral.App, format string) {
	if len(apps) == 0 {
		fmt.Printf("No apps are found\n")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "App Name", "Subdomain", "Resource Name", "Location", "SKU", "Template"})

	for i, item := range apps {
		t.AppendRow([]interface{}{i + 1, *item.AppProperties.ApplicationID, *item.AppProperties.DisplayName,
			*item.AppProperties.Subdomain, *item.Name, *item.Location, item.Sku.Name, *item.AppProperties.Template})
	}
	util.RenderTable(t, format, false)
}
