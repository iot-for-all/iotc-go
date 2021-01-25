package apitokens

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/roles"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the API Tokens list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of API tokens in an application",
	Long: `Get the list of API tokens in an application.
The token value will never be returned for security reasons.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		app, err := cmd.Flags().GetString("app")
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		top, err := cmd.Flags().GetInt("top")
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

		// get the list of API Tokens
		res, err := c.Operations.APITokensList(operations.NewAPITokensListParams())
		if err != nil {
			return err
		}

		// get the list of roles
		spin.Suffix = " Downloading roles"
		roles, err := roles.GetRolesLookupTable(c, app)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Role", "Expiry"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, res.Payload.Value, roles, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := res.Payload.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v actions, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var tokens models.APITokenCollection
			if err := tokens.UnmarshalBinary(body); err != nil {
				return err
			}

			numItem, limitReached, moreRowsExist = addTableRows(t, tokens.Value, roles, numItem, top)

			nextLink = tokens.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	apiTokensCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func addTableRows(t table.Writer, tokens []*models.APIToken, roles map[string]string, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range tokens {
		rolesStr := ""
		for _, r := range item.Roles {
			if len(rolesStr) > 0 {
				rolesStr += ", "
			}
			rolesStr += roles[*r.Role]
		}

		t.AppendRow([]interface{}{numItem, item.ID, rolesStr, item.Expiry})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(tokens) != i+1
			break
		}

		numItem++
	}
	return numItem, limitReached, moreRowsExist
}
