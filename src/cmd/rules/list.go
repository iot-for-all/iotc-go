package rules

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the rules list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the rules in an application",
	Long:  `List all the rules in an application`,
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
		spin := util.NewSpinner(" Downloading rules ...")

		// get the list of rules
		res, err := c.Operations.RulesList(operations.NewRulesListParams())
		if err != nil {
			return err
		}

		if len(res.Payload.Value) == 0 {
			fmt.Printf("No rules found in '%s' app\n", app)
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Display Name"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, res.Payload.Value, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := res.Payload.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v rules, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var rc models.RuleCollection
			if err := rc.UnmarshalBinary(body); err != nil {
				return err
			}
			numItem, limitReached, moreRowsExist = addTableRows(t, rc.Value, numItem, top)

			nextLink = rc.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	rulesCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func GetRulesLookupTable(client *client.AzureIoTCentral, app string) (map[string]string, error) {
	rules := make(map[string]string)

	res, err := client.Operations.RulesList(operations.NewRulesListParams())
	if err != nil {
		return rules, err
	}
	for _, item := range res.Payload.Value {
		rules[item.ID] = item.DisplayName
	}

	// see if there are more results beyond first page
	nextLink := res.Payload.NextLink
	for {
		if len(nextLink) == 0 {
			break
		}

		body, err := util.GetContent(app, nextLink)
		if err != nil {
			return rules, err
		}

		var rc models.RuleCollection
		if err := rc.UnmarshalBinary(body); err != nil {
			return rules, err
		}
		for _, item := range rc.Value {
			rules[item.ID] = item.DisplayName
		}

		nextLink = rc.NextLink
	}

	return rules, nil
}

func addTableRows(t table.Writer, rules []*models.Rule, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range rules {
		t.AppendRow([]interface{}{numItem, item.ID, item.DisplayName})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(rules) != i+1
			break
		}
		numItem++
	}
	return numItem, limitReached, moreRowsExist
}
