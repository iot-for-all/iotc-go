package users

import (
	"bytes"
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// setCmd represents the users set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Add a new user to an application",
	Long: `Add a new user to an application`,
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
		userType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		tenantID, err := cmd.Flags().GetString("tenantID")
		if err != nil {
			return err
		}
		objectID, err := cmd.Flags().GetString("objectID")
		if err != nil {
			return err
		}

		switch userType{
		case "EmailUser":
			if email == "" {
				return errors.New("email is required for EmailUser")
			}
		case "ServicePrincipalUser":
			if tenantID == "" || objectID == "" {
				return errors.New("tenantID and objectID are required for ServicePrincipalUser")
			}
		default:
			return errors.New("unknown user type, only EmailUser and ServicePrincipalUser are allowed")
		}

		//change this to send email and service principal
		user := &util.CustomUser {
			ID:         id,
			Type:       userType,
		}
		user.Roles = make([]*util.CustomRolesItems, 1)
		user.Roles[0] = &util.CustomRolesItems{Role: role}
		if userType == "EmailUser" {
			user.Email = &email
		} else {
			user.TenantID = &tenantID
			user.ObjectID = &objectID
		}

		// Set the User
		/*p := operations.NewUsersSetParams()
		p.UserID = id
		user := &models.User{
			ID:         id,
			Type:       "EmailUser",
		}
		p.Body = user
		user.Roles = make([]*models.PermissionRolesItems0, 1)
		user.Roles[0] = &models.PermissionRolesItems0{Role: &role}
		_, err = c.Operations.UsersSet(p)
		if err != nil {
			return err
		}*/

		url, err := client.GetURL(app)
		if err != nil {
			return err
		}
		url += "/users/" + id

		request, err := json.Marshal(user)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Uploading user ...")

		resp, err := util.PutContent(app, url, request)
		if err != nil {
			return err
		}

		spin.Stop()

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, resp, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", prettyJSON.String())

		return nil
	},
}

func init() {
	usersCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	setCmd.MarkFlagRequired("app")
	setCmd.Flags().StringP("id", "", "", "user ID of the user")
	setCmd.MarkFlagRequired("id")
	setCmd.Flags().StringP("role", "", "", "role ID to be assigned for this user")
	setCmd.MarkFlagRequired("role")
	setCmd.Flags().StringP("type", "", "", "type of the user (EmailUser or ServicePrincipalUser)")
	setCmd.MarkFlagRequired("type")
	setCmd.Flags().StringP("email", "", "", "email address of the EmailUser")
	setCmd.Flags().StringP("tenantID", "", "", "tenant ID of the ServicePrincipalUser")
	setCmd.Flags().StringP("objectID", "", "", "object ID of the ServicePrincipalUser")
}
