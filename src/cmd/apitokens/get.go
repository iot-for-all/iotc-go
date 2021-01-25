package apitokens

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/roles"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the API Tokens get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an API token by ID",
	Long: `Get an API token by ID`,
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
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading API Tokens ...")

		// Get the API Token
		p := operations.NewAPITokensGetParams()
		p.TokenID = id
		res, err := c.Operations.APITokensGet(p)
		if err != nil {
			return err
		}

		// get the list of roles
		spin.Suffix = " Downloading roles"
		roles, err := roles.GetRolesLookupTable(c, app)
		if err != nil {
			return err
		}

		spin.Stop()
		printAPIToken(res.Payload, roles, app, format)

		return nil
	},
}

func init() {
	apiTokensCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCmd.MarkFlagRequired("app")
	getCmd.Flags().StringP("id", "", "", "unique ID of the API token")
	getCmd.MarkFlagRequired("id")
	getCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

// printAPIToken prints the API Token
func printAPIToken(apiToken *models.APIToken, roles map[string]string, app string, format string) {
	rolesStr := ""
	for _, r := range apiToken.Roles{
		if len(rolesStr) > 0 {
			rolesStr += ", "
		}
		rolesStr += fmt.Sprintf("%s (%s)", roles[*r.Role], *r.Role)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", apiToken.ID})
	t.AppendRow([]interface{}{"Role", rolesStr})
	t.AppendRow([]interface{}{"Expiry", apiToken.Expiry})

	util.RenderTable(t, format, false)
}