package actions

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "Get the actions used in rules",
	Long: `Actions are triggered when a rule evaluates to true. Get the actions used in all the rules.`,
}

func init() {
	cmd.RootCmd.AddCommand(actionsCmd)
}