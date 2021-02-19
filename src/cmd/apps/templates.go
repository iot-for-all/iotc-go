package apps

import (
	"context"
	"fmt"
	"os"

	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listTemplatesCmd represents the app templates list command
var listTemplatesCmd = &cobra.Command{
	Use:   "listTemplates",
	Short: "Get the list of all the app templates",
	Long:  `Get the list of all the app templates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		subscription, err := cmd.Flags().GetString("subscription")
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
		spin.Suffix = " Downloading app templates list..."
		templates, err := client.ListTemplates(context.Background())
		if err != nil {
			return err
		}

		// write out the table
		spin.Stop()
		printTemplatesTable(templates.Values(), format)

		return nil
	},
}

func init() {
	appsCmd.AddCommand(listTemplatesCmd)

	listTemplatesCmd.Flags().StringP("subscription", "s", "", "Azure subscription ID")
	listTemplatesCmd.MarkFlagRequired("subscription")
	listTemplatesCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

// printTemplatesTable prints the app templates list as a table
func printTemplatesTable(templates []iotcentral.AppTemplate, format string) {
	if len(templates) == 0 {
		fmt.Printf("No app templates are found\n")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "Title", "Description", "Version", "Order", "Template Name"})

	for i, item := range templates {
		t.AppendRow([]interface{}{i + 1, *item.ManifestID, *item.Title, *item.Description, *item.ManifestVersion, *item.Order, *item.Name})
	}
	util.RenderTable(t, format, false)
}
