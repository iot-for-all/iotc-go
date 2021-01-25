package jobs

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicegroups"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the devices list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the jobs in an application",
	Long:  `List all the jobs in an application`,
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
		spin := util.NewSpinner(" Downloading devices ...")

		// get the list of jobs
		res, err := c.Operations.JobsList(operations.NewJobsListParams())
		if err != nil {
			return err
		}

		if len(res.Payload.Value) == 0 {
			fmt.Printf("No jobs found in '%s' app\n", app)
			return nil
		}

		// get all device templates look up table so that we can print the template names

		// get all device templates look up table so that we can print the template names
		spin.Suffix = " Getting device groups"
		deviceGroups, err := devicegroups.GetDeviceGroupLookupTable(c, app)
		if err != nil {
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "ID", "Name", "Description", "Group", "Status"})

		numItem := 1
		limitReached := false
		moreRowsExist := false
		numItem, limitReached, moreRowsExist = addTableRows(t, res.Payload.Value, deviceGroups, numItem, top)

		// loop through and download all the rows one page at a time
		nextLink := res.Payload.NextLink
		for {
			if len(nextLink) == 0 || limitReached {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v jobs, getting more...", numItem-1)
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var jc models.JobCollection
			if err := jc.UnmarshalBinary(body); err != nil {
				return err
			}
			numItem, limitReached, moreRowsExist = addTableRows(t, jc.Value, deviceGroups, numItem, top)

			nextLink = jc.NextLink
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, moreRowsExist || len(nextLink) != 0)

		return nil
	},
}

func init() {
	jobsCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	listCmd.MarkFlagRequired("app")
	listCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	listCmd.Flags().IntP("top", "", config.Config.MaxRows, "list only top N rows")
}

func addTableRows(t table.Writer, jobs []*models.Job, deviceGroups map[string]string, numItem int, top int) (int, bool, bool) {
	var limitReached = false
	var moreRowsExist = false
	for i, item := range jobs {
		t.AppendRow([]interface{}{numItem, item.ID, item.DisplayName, item.Description, deviceGroups[*item.Group], item.Status})
		if numItem == top {
			limitReached = true
			moreRowsExist = len(jobs) != i+1
			break
		}
		numItem++
	}
	return numItem, limitReached, moreRowsExist
}
