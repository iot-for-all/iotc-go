package rules

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// rulesCmd represents the rules command
var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "List rules within your application",
	Long: `List rules within your application.`,
}

func init() {
	cmd.RootCmd.AddCommand(rulesCmd)
}