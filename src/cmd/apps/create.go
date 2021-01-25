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

// createCmd represents the create app command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create or update an app in the given subscription",
	Long:  `Create or update an app in the given subscription.`,
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
		location, err := cmd.Flags().GetString("location")
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

		// check for app name availability
		spin.Suffix = " Checking name availability..."
		nameAvailability := iotcentral.OperationInputs{
			Name: &appName,
		}
		result, err := client.CheckNameAvailability(context.Background(), nameAvailability)
		if err != nil {
			return err
		}
		if !*result.NameAvailable {
			return errors.New(fmt.Sprintf("'%s' app is already in use, use a different name", appName))
		}

		// check for subdomain name availability
		spin.Suffix = " Checking subdomain name availability..."
		var appType string = "IoTApps"
		subdomainAvailability := iotcentral.OperationInputs{
			Name: &subdomain,
			Type: &appType,
		}
		sdResult, err := client.CheckSubdomainAvailability(context.Background(), subdomainAvailability)
		if err != nil {
			return err
		}
		if !*sdResult.NameAvailable {
			return errors.New(fmt.Sprintf("'%s' subdomain is already in use, use a different name", subdomain))
		}

		// create app
		spin.Suffix = " Creating IoT Central app..."
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
		app := iotcentral.App{
			AppProperties: &iotcentral.AppProperties{
				DisplayName: &appName,
				Subdomain:   &subdomain,
			},
			Sku: &iotcentral.AppSkuInfo{
				Name: appSKU,
			},
			Name:     &appName,
			Location: &location,
		}

		// create the app
		_, err = client.CreateOrUpdate(context.Background(), resourceGroup, resourceName, app)
		if err != nil {
			return err
		}
		//fmt.Println(createResult.Response())

		// stop the spinner
		spin.Stop()

		return nil
	},
}

func init() {
	appsCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("subscription", "s", "", "Azure subscription ID")
	createCmd.MarkFlagRequired("subscription")
	createCmd.Flags().StringP("resourceGroup", "", "", "the resource group ID of the application")
	createCmd.MarkFlagRequired("resourceGroup")
	createCmd.Flags().StringP("resourceName", "a", "", "the name of the application resource")
	createCmd.MarkFlagRequired("resourceName")
	createCmd.Flags().StringP("appName", "", "", "the name of the application")
	createCmd.MarkFlagRequired("appName")
	createCmd.Flags().StringP("subdomain", "", "", "the name application subdomain")
	createCmd.MarkFlagRequired("subdomain")
	createCmd.Flags().StringP("sku", "", "", "the IoT Central sku ST0, ST1 or ST2")
	createCmd.MarkFlagRequired("sku")
	createCmd.Flags().StringP("location", "", "", "the location of the application")
	createCmd.MarkFlagRequired("location")
}
