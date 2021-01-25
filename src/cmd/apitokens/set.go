package apitokens

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"time"
)

// setCmd represents the API Tokens set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a new API token in the application to use in the IoT Central public API",
	Long: `Create a new API token in the application to use in the IoT Central public API.
The token value will be returned in the response, and won't be returned again in subsequent requests.`,
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
		role, err := cmd.Flags().GetString("role")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// Set the API Token
		p := operations.NewAPITokensSetParams()
		p.TokenID = id
		token := &models.APIToken {
			ID: id,
			Expiry: strfmt.DateTime(time.Now().AddDate(1, 0, 0)),
		}
		p.Body = token
		token.Roles = make([]*models.PermissionRolesItems0, 1)
		token.Roles[0] = &models.PermissionRolesItems0{Role: &role}

		res, err := c.Operations.APITokensSet(p)
		if err != nil {
			return err
		}

		fmt.Printf("Token: %s\n", res.Payload.Token)

		return nil
	},
}

func init() {
	apiTokensCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setCmd.MarkFlagRequired("app")
	setCmd.Flags().StringP("id", "", "", "unique ID of the API token")
	setCmd.MarkFlagRequired("id")
	setCmd.Flags().StringP("role", "", "", "role ID that specify the permissions to access the application")
	setCmd.MarkFlagRequired("role")
}
