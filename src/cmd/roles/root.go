package roles

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// rolesCmd represents the roles command
var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "List roles within your application",
	Long: `List roles within your application.`,
}

func init() {
	cmd.RootCmd.AddCommand(rolesCmd)
}