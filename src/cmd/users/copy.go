package users

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// copyCmd represents the users copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy users from one application to another",
	Long: `Copy users from one application to another.

Each user belongs to a role and the role assignment is done using role ID.
The role ID is assigned during the role creation. Currently IoT Central does
not expose a way to create roles or get role permissions. There is no way
to verify that the same roles exist in both apps except for the builtin
default roles 'Administrator', 'Builder', 'Operator'. So, only the users
belonging to these default roles are copied and the rest of the users or
user role assignments are not copied.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		srcApp, err := cmd.Flags().GetString("srcApp")
		if err != nil {
			return err
		}
		destApp, err := cmd.Flags().GetString("destApp")
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading users ...")

		// get all users from src and dest apps
		srcUsers, err := getUsers(srcApp, spin)
		if err != nil {
			return err
		}
		destUsers, err := getUsers(destApp, spin)
		if err != nil {
			return err
		}

		var defaultRoles = []string{
			"ca310b8d-2f4a-44e0-a36e-957c202cd8d4", // Administrator
			"344138e9-8de4-4497-8c54-5237e96d6aaf", // Builder
			"ae2c9854-393b-4f97-8c42-479d70ce626e", // Operator
		}

		// copy all users belonging to default roles
		var addedUsers = 0
		var updatedUsers = 0
		var sameUsers = 0
		var skippedUsers = 0
		for _, srcUser := range srcUsers {
			var newUser = util.CustomUser{
				ID:               srcUser.ID,
				Type:             srcUser.Type,
				Email:            srcUser.Email,
				ObjectID:         srcUser.ObjectID,
				TenantID:         srcUser.TenantID,
				CustomPermission: util.CustomPermission{},
			}

			// copy only if the role assignments for default role
			for _, userRole := range srcUser.Roles {
				for _, defaultRole := range defaultRoles {
					if defaultRole == userRole.Role {
						newUser.CustomPermission.Roles = append(newUser.CustomPermission.Roles, &util.CustomRolesItems{Role: defaultRole})
						break
					}
				}
			}

			// add only if we have role assignments
			// i.e. default roles from source app user and existing role assignments from dest app
			if len(newUser.Roles) > 0 {

				var foundExistingUser = false
				var newRolesAdded = true
				for _, destUser := range destUsers {
					if newUser.Type == "EmailUser" && destUser.Type == "EmailUser" && *newUser.Email == *destUser.Email {
						newUser.ID = destUser.ID
						newUser.Roles, newRolesAdded = mergeRoles(newUser.Roles, destUser.Roles)
						foundExistingUser = true
						break
					} else if newUser.Type == "ServicePrincipalUser" && destUser.Type == "ServicePrincipalUser" && *newUser.TenantID == *destUser.TenantID && *newUser.ObjectID == *destUser.ObjectID {
						newUser.ID = destUser.ID
						newUser.Roles, newRolesAdded = mergeRoles(newUser.Roles, destUser.Roles)
						foundExistingUser = true
						break
					}
				}

				if newRolesAdded {
					url, err := client.GetURL(destApp)
					if err != nil {
						return err
					}
					url += "/users/" + newUser.ID

					request, err := json.Marshal(newUser)
					if err != nil {
						return err
					}

					if srcUser.Type == "EmailUser" {
						spin.Suffix = fmt.Sprintf(" Uploading '%s' to '%s'", *newUser.Email, destApp)
					} else {
						spin.Suffix = fmt.Sprintf(" Uploading '%s' to '%s'", *newUser.ObjectID, destApp)
					}

					_, err = util.PutContent(destApp, url, request)
					if err != nil {
						return err
					}

					if foundExistingUser {
						updatedUsers++
					} else {
						addedUsers++
					}
				} else {
					sameUsers++
				}
			} else {
				skippedUsers++
			}
		}

		spin.Stop()
		fmt.Printf("%v new user(s) added to '%s' app\n", addedUsers, destApp)
		fmt.Printf("%v existing user(s) updated in '%s' app\n", updatedUsers, destApp)
		fmt.Printf("%v user(s) already exist with same role assignments, so skipped them in '%s' app\n", sameUsers, destApp)
		if skippedUsers > 0 {
			fmt.Printf("%v user(s) are not copied as they have only custom role assignments\n", skippedUsers)
			fmt.Printf("See 'users copy --help' for more info\n")
		}

		return nil
	},
}

func init() {
	usersCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringP("srcApp", "s", "", "source application to copy from")
	copyCmd.MarkFlagRequired("srcApp")
	copyCmd.Flags().StringP("destApp", "d", "", "destination application to copy into")
	copyCmd.MarkFlagRequired("destApp")
}

func getUsers(app string, spin *spinner.Spinner) ([]util.CustomUser, error) {

	var users []util.CustomUser
	spin.Suffix = fmt.Sprintf(" Downloading users from '%s'...", app)

	// get the list of users
	url, err := client.GetURL(app)
	if err != nil {
		return users, err
	}

	url += "/users"
	body, err := util.GetContent(app, url)
	if err != nil {
		return users, err
	}
	var uc util.CustomUserCollection
	if err := uc.UnmarshalBinary(body); err != nil {
		return users, err
	}

	users = append(users, uc.Value...)

	// loop through and download all the rows one page at a time
	nextLink := uc.NextLink
	numItem := len(users)
	for {
		if len(nextLink) == 0 {
			break
		}

		spin.Suffix = fmt.Sprintf(" Downloaded %v users from %s, getting more...", numItem, app)
		body, err := util.GetContent(app, nextLink)
		if err != nil {
			return users, err
		}

		var nuc util.CustomUserCollection
		if err := nuc.UnmarshalBinary(body); err != nil {
			return users, err
		}

		users = append(users, uc.Value...)
		nextLink = uc.NextLink
	}

	return users, nil
}

// mergeRoles add existing roles into source roles and return new set of roles
func mergeRoles(srcRoles []*util.CustomRolesItems, existingDestRoles []*util.CustomRolesItems) ([]*util.CustomRolesItems, bool) {
	var roles = srcRoles
	var rolesAreEqual = false
	if len(srcRoles) == len(existingDestRoles) {
		rolesAreEqual = true
		for _, sr := range srcRoles {
			var found = false
			for _, er := range existingDestRoles {
				if sr.Role == er.Role {
					found = true
					break
				}
			}
			if !found {
				rolesAreEqual = false
				break
			}
		}
	}

	if rolesAreEqual {
		return roles, false
	}

	for _, er := range srcRoles {
		var found = false
		for _, nr := range existingDestRoles {
			if nr.Role == er.Role {
				found = true
				break
			}
		}
		if !found {
			roles = append(roles, &util.CustomRolesItems{Role: er.Role})
		}
	}

	return roles, true
}
