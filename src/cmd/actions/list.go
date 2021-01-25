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

// listCmd represents the actions list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the actions in an application",
	Long:  `Get the actions in an application`,
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
		spin := util.NewSpinner(" Downloading actions ...")

		// get the list of actions
		res, err := c.Operations.ActionsList(operations.NewActionsListParams())
		if err != nil {
			return err
		}

		if len(res.Payload.Value()) == 0 {
			fmt.Printf("No actions found in '%s' app\n", app)
			return nil
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Display Name", "Type", "Rules"})

		// get the list of rules
		rules, err := rules.GetRulesLookupTable(c, app)
		if err != nil {
			return err
		}

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, res.Payload.Value(), rules, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := res.Payload.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v actions, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var actions models.ActionCollection
			if err := actions.UnmarshalBinary(body); err != nil {
				return err
			}

			numItem, limitReached, moreRowsExist = addTableRows(t, actions.Value(), rules, numItem, top)

			nextLink = actions.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	actionsCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func addTableRows(t table.Writer, actions []models.Action, rules map[string]string, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range actions {
		rulesStr := ""
		for _, r := range item.Rules() {
			rulesStr = rulesStr + rules[r] + " "
		}
		t.AppendRow([]interface{}{numItem, item.ID(), item.DisplayName(), item.Type(), rulesStr})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(actions) != i+1
			break
		}

		numItem++
	}
	return numItem, limitReached, moreRowsExist
}
