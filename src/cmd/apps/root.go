package apps

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Create, update and delete  IoT Central applications",
	Long: `Create, update and delete  IoT Central applications.
This makes calls to Azure Resource Manager (ARM) to accomplish these changes.
You need to login to your Azure account using Azure CLI 'az login'
Download the Azure CLI from: https://docs.microsoft.com/en-us/cli/azure/
`,
}

func init() {
	cmd.RootCmd.AddCommand(appsCmd)
}