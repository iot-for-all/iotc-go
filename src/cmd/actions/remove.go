package actions

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"github.com/spf13/cobra"
)

// removeCmd represents the actions remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete an action",
	Long: `Delete an action from an application`,
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

		// remove the action
		p := operations.NewActionsRemoveParams()
		p.ActionID = id
		_, err = c.Operations.ActionsRemove(p)
		if err != nil {
			return err
		}

		//fmt.Printf("Res: %v\n", res.Error())

		return nil
	},
}

func init() {
	actionsCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	removeCmd.MarkFlagRequired("app")
	removeCmd.Flags().StringP("id", "", "", "unique ID of the action")
	removeCmd.MarkFlagRequired("id")
}
