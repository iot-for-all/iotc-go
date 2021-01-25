package actions

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/rules"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// gtCmd represents the actions get command
var gtCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an Action by ID",
	Long: `Get an Action by ID`,
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
		spin := util.NewSpinner(" Downloading API Tokens ...")

		// get the action with given ID
		p := operations.NewActionsGetParams()
		p.ActionID = id
		res, err := c.Operations.ActionsGet(p)
		if err != nil {
			return err
		}

		// get the list of rules
		spin.Suffix = " Downloading rules"
		rules, err := rules.GetRulesLookupTable(c, app)
		if err != nil {
			return err
		}

		spin.Stop()
		printAction(res.Payload, rules, app, format)

		return nil
	},
}

func init() {
	actionsCmd.AddCommand(gtCmd)

	gtCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	gtCmd.MarkFlagRequired("app")
	gtCmd.Flags().StringP("id", "", "", "unique ID for the Action")
	gtCmd.MarkFlagRequired("id")
	gtCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
}

// printAction prints the action
func printAction(action models.Action, rules map[string]string, app string, format string) {
	if action == nil {
		fmt.Printf("No such actions found in '%s' app\n", app)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", action.ID()})
	t.AppendRow([]interface{}{"DisplayName", action.DisplayName()})
	t.AppendRow([]interface{}{"Type", action.Type()})

	for _, item := range action.Rules(){
		t.AppendRow([]interface{}{"Rule", fmt.Sprintf("%s (%s)", rules[item], item)})
	}

	util.RenderTable(t, format, false)
}