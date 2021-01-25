package users

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/cmd/roles"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the users get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an user by ID",
	Long: `Get an user by ID`,
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
		spin := util.NewSpinner(" Downloading user ...")

		// Get the user
		//p := operations.NewUsersGetParams()
		//p.UserID = id
		//res, err := c.Operations.UsersGet(p)
		//if err != nil {
		//	return err
		//}
		url, err := client.GetURL(app)
		if err != nil {
			return err
		}

		url += "/users/" + id
		body, err := util.GetContent(app, url)
		if err != nil {
			return err
		}
		var user util.CustomUser
		if err := json.Unmarshal(body, &user); err != nil {
			return err
		}

		// get the list of roles
		spin.Suffix = " Downloading roles"
		roles, err := roles.GetRolesLookupTable(c, app)
		if err != nil {
			return err
		}

		spin.Stop()
		printUser(user, roles, app, format)

		return nil
	},
}

func init() {
	usersCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCmd.MarkFlagRequired("app")
	getCmd.Flags().StringP("id", "", "", "unique ID for the user")
	getCmd.MarkFlagRequired("id")
	getCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}


// printUser prints the user
func printUser(user util.CustomUser, roles map[string]string, app string, format string) {
	rolesStr := ""
	for _, r := range user.Roles {
		if len(rolesStr) > 0 {
			rolesStr += ", "
		}
		rolesStr += roles[r.Role]
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", user.ID})
	t.AppendRow([]interface{}{"Type", user.Type})
	t.AppendRow([]interface{}{"Role", rolesStr})
	if user.Type == "EmailUser" {
		email := ""
		if user.Email != nil {
			email = *user.Email
		}
		t.AppendRow([]interface{}{"Email", email})
	} else if user.Type == "ServicePrincipalUser" {
		tenantID := ""
		if user.TenantID != nil {
			tenantID = *user.TenantID
		}
		objectID := ""
		if user.ObjectID != nil {
			objectID = *user.ObjectID
		}
		t.AppendRow([]interface{}{"Tenant ID", tenantID})
		t.AppendRow([]interface{}{"Object ID", objectID})
	}

	util.RenderTable(t, format, false)
}

// printUser prints the user
func printUserOld(user map[string]interface{}, roles map[string]string, app string, format string) {
	id := user["id"].(string)
	userType := user["type"].(string)
	email := ""
	tenantID := ""
	objectID := ""
	rolesStr := ""

	userRoles := user["roles"].([]interface{})
	for _, r := range userRoles {
		roleID := r.(map[string]interface{})["role"].(string)
		if len(rolesStr) > 0 {
			rolesStr += ", "
		}
		rolesStr += roles[roleID]
	}

	switch userType {
	case "EmailUser":
		email = user["email"].(string)
	case "ServicePrincipalUser":
		tenantID = user["tenantId"].(string)
		objectID = user["objectId"].(string)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", id})
	t.AppendRow([]interface{}{"Type", userType})
	t.AppendRow([]interface{}{"Role", rolesStr})
	t.AppendRow([]interface{}{"Email", email})
	t.AppendRow([]interface{}{"Tenant ID", tenantID})
	t.AppendRow([]interface{}{"Object ID", objectID})

	util.RenderTable(t, format, false)
}