package apitokens

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// apiTokensCmd represents the apiTokens command
var apiTokensCmd = &cobra.Command{
	Use:   "apiTokens",
	Short: "Create, read, delete access tokens used to interact with the IoT Central public APIs",
	Long: `Create, read, delete access tokens used to interact with the IoT Central public APIs.`,
}

func init() {
	cmd.RootCmd.AddCommand(apiTokensCmd)
}