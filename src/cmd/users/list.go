package users

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/cmd/roles"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the users list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the users in an application",
	Long:  `List all the users in an application`,
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
		spin := util.NewSpinner(" Downloading users ...")

		// get the list of users
		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/users"
		body, err := util.GetContent(app, url)
		if err != nil {
			return err
		}
		var uc util.CustomUserCollection
		if err := uc.UnmarshalBinary(body); err != nil {
			return err
		}

		// get the list of roles
		spin.Suffix = " Downloading roles ..."
		roles, err := roles.GetRolesLookupTable(c, app)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Type", "Roles", "Email", "Tenant ID", "Object ID"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, uc.Value, roles, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := uc.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v users, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var nuc util.CustomUserCollection
			if err := nuc.UnmarshalBinary(body); err != nil {
				return err
			}

			numItem, limitReached, moreRowsExist = addTableRows(t, nuc.Value, roles, numItem, top)

			nextLink = uc.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	usersCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func addTableRows(t table.Writer, users []util.CustomUser, roles map[string]string, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, user := range users {
		rolesStr := ""
		for _, r := range user.Roles {
			if len(rolesStr) > 0{
				rolesStr += ", "
			}
			rolesStr += roles[r.Role]
		}

		email := ""
		if user.Email != nil {
			email = *user.Email
		}
		tenantID := ""
		if user.TenantID != nil {
			tenantID = *user.TenantID
		}
		objectID := ""
		if user.ObjectID != nil {
			objectID = *user.ObjectID
		}
		t.AppendRow([]interface{}{numItem, user.ID, user.Type, rolesStr, email, tenantID, objectID})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(users) != i+1
			break
		}
		numItem++
	}
	return numItem, limitReached, moreRowsExist
}
