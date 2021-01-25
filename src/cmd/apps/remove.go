package apps

import (
	"com.azure.iot/iotcentral/iotcgo/util"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/spf13/cobra"
)

// removeCmd represents the apps remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete an IoT Central application",
	Long: `Delete an IoT Central application.
You need to specify the application resource name to remove. You can get it
using the 'apps list' command.
`,
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

		// delete the app from the subscription
		spin.Suffix = " Deleting IoT Central app..."
		_, err = client.Delete(context.Background(), resourceGroup, resourceName)
		if err != nil {
			return err
		}

		// write out the table
		spin.Stop()

		return nil
	},
}

func init() {
	appsCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("subscription", "", "", "the Azure subscription ID of the application")
	removeCmd.MarkFlagRequired("subscription")
	removeCmd.Flags().StringP("resourceGroup", "", "", "the resource group ID of the application")
	removeCmd.MarkFlagRequired("resourceGroup")
	removeCmd.Flags().StringP("resourceName", "a", "", "the name of the application resource")
	removeCmd.MarkFlagRequired("resourceName")
}
