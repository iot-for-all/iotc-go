package users

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Add, update, copy (between apps) and remove users within your application",
	Long:  `Add, update, copy (between apps) and remove users within your application.`,
}

func init() {
	cmd.RootCmd.AddCommand(usersCmd)
}
