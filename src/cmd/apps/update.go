package apps

import (
	"com.azure.iot/iotcentral/iotcgo/util"
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/spf13/cobra"
	"strings"
)

// updateCmd represents the update app command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an app in the given subscription",
	Long:  `Update an app in the given subscription.`,
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
		appName, err := cmd.Flags().GetString("appName")
		if err != nil {
			return err
		}
		subdomain, err := cmd.Flags().GetString("subdomain")
		if err != nil {
			return err
		}
		sku, err := cmd.Flags().GetString("sku")
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

		// create app
		spin.Suffix = " Updating IoT Central app..."

		var patch iotcentral.AppPatch
		var appProperties iotcentral.AppProperties
		patch.AppProperties = &appProperties
		if len(appName) > 0 {
			appProperties.DisplayName = &appName
		}
		if len(subdomain) > 0 {
			appProperties.Subdomain = &subdomain
		}
		if len(sku) > 0 {
			var appSKU iotcentral.AppSku
			switch strings.ToLower(sku) {
			case "st0":
				appSKU = iotcentral.ST0
			case "st1":
				appSKU = iotcentral.ST1
			case "st2":
				appSKU = iotcentral.ST2
			default:
				return errors.New(fmt.Sprintf("Unknown SKU '%s'. Only %s, %s, %s are allowed", sku, iotcentral.ST0, iotcentral.ST1, iotcentral.ST2))
			}
			patch.Sku = &iotcentral.AppSkuInfo{
				Name: appSKU,
			}
		}

		// update the app
		_, err = client.Update(context.Background(), resourceGroup, resourceName, patch)
		if err != nil {
			return err
		}

		// stop the spinner
		spin.Stop()

		return nil
	},
}

func init() {
	appsCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("subscription", "s", "", "Azure subscription ID")
	updateCmd.MarkFlagRequired("subscription")
	updateCmd.Flags().StringP("resourceGroup", "", "", "the resource group ID of the application")
	updateCmd.MarkFlagRequired("resourceGroup")
	updateCmd.Flags().StringP("resourceName", "a", "", "the name of the application resource")
	updateCmd.MarkFlagRequired("resourceName")
	updateCmd.Flags().StringP("appName", "", "", "the updated name of the application")
	updateCmd.Flags().StringP("subdomain", "", "", "the updated name application subdomain")
	updateCmd.Flags().StringP("sku", "", "", "the updated IoT Central sku ST0, ST1 or ST2")
}
