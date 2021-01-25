package apitokens

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"github.com/spf13/cobra"
)

// removeCmd represents the API Tokens remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete an API token",
	Long: `Delete an API token from an application`,
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
		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// remove the API Token
		p := operations.NewAPITokensRemoveParams()
		p.TokenID = id
		_, err = c.Operations.APITokensRemove(p)
		if err != nil {
			return err
		}

		//fmt.Printf("Res: %v\n", res.Error())

		return nil
	},
}

func init() {
	apiTokensCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	removeCmd.MarkFlagRequired("app")
	removeCmd.Flags().StringP("id", "", "", "unique ID of the API token")
	removeCmd.MarkFlagRequired("id")
}
