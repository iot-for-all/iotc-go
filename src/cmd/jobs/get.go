package jobs

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicegroups"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

// getCmd represents the jobs get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a job by ID",
	Long:  `Get a job by ID.
If the '--outputFile' is specified, JSON representation of the given job
is written to the file and rest of the parameters are ignored.
This JSON file can be easily edited as a template to be used in the
'set' command.`,
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
		showCompleted, err := cmd.Flags().GetBool("showCompleted")
		if err != nil {
			return err
		}
		showFailed, err := cmd.Flags().GetBool("showFailed")
		if err != nil {
			return err
		}
		showPending, err := cmd.Flags().GetBool("showPending")
		if err != nil {
			return err
		}
		outputFile, err := cmd.Flags().GetString("outputFile")
		if err != nil {
			return err
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading job ...")

		// if the output file is given, write the contents into a file
		if len(outputFile) > 0 {
			url, err := client.GetURL(app)
			if err != nil {
				return err
			}

			url = url + "/jobs/" + id
			body, err := util.GetIndentedJSONContent(app, url) // get the
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(outputFile, body, 0755)
			if err != nil {
				return err
			}

			spin.Stop()
			fmt.Printf("Job properties are written into the file: %s\n", outputFile)
			return nil
		}

		// Get the job
		p := operations.NewJobsGetParams()
		p.JobID = id
		res, err := c.Operations.JobsGet(p)
		if err != nil {
			return err
		}

		// get the list of device groups
		spin.Suffix = " Downloading device group"
		deviceGroups, err := devicegroups.GetDeviceGroupLookupTable(c, app) // TODO Get a single device group
		if err != nil {
			return err
		}

		// get the list of devices
		spin.Suffix = " Downloading devices"
		jdp := operations.NewJobsGetDevicesParams()
		jdp.JobID = id
		devices, err := c.Operations.JobsGetDevices(jdp)
		if err != nil {
			return err
		}

		// loop through and download all the rows one page at a time
		nextLink := devices.Payload.NextLink
		var allDevices []*models.JobDeviceStatus
		allDevices = append(allDevices, devices.Payload.Value...)
		for {
			if len(nextLink) == 0 {
				break
			}

			spin.Suffix = fmt.Sprintf(" Downloaded %v devices, getting more...", len(allDevices))
			body, err := util.GetContent(app, nextLink)
			if err != nil {
				return err
			}

			var jdc models.JobDeviceStatusCollection
			if err := jdc.UnmarshalBinary(body); err != nil {
				return err
			}

			allDevices = append(allDevices, jdc.Value...)
			nextLink = jdc.NextLink
		}

		spin.Stop()
		printJob(res.Payload, deviceGroups, allDevices, app, format, showCompleted, showFailed, showPending)

		return nil
	},
}

func init() {
	jobsCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getCmd.MarkFlagRequired("app")
	getCmd.Flags().StringP("id", "", "", "unique ID for the job")
	getCmd.MarkFlagRequired("id")
	getCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	getCmd.Flags().StringP("outputFile", "", "", "dump the job JSON to the given file")
	getCmd.Flags().BoolP("showCompleted", "", false, "show all devices that the job has completed updating")
	getCmd.Flags().BoolP("showFailed", "", false, "show all devices that the job has failed to update")
	getCmd.Flags().BoolP("showPending", "", false, "show all devices that the job is pending to update")
}

// printJob prints the job
func printJob(job *models.Job, deviceGroups map[string]string, devices []*models.JobDeviceStatus, app string,
	format string, showCompleted bool, showFailed bool, showPending bool) {

	// print job header
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow([]interface{}{"ID", job.ID})
	t.AppendRow([]interface{}{"Name", job.DisplayName})
	t.AppendRow([]interface{}{"Description", job.Description})
	t.AppendRow([]interface{}{"Group", deviceGroups[*job.Group]})
	t.AppendRow([]interface{}{"Status", job.Status})
	util.RenderTable(t, format, false)

	// print job parameters i.e. what the job is changing
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Type", "Target", "Path", "Value"})
	for _, jd := range job.Data() {
		var val string = ""
		if jd.Value() != nil {
			val = fmt.Sprintf("%v", jd.Value())
		}
		t.AppendRow([]interface{}{jd.Type(), *jd.Target(), *jd.Path(), val})
	}
	util.RenderTable(t, format, false)

	// print job stats
	var completed int = 0
	var failed int = 0
	var pending int = 0
	for _, device := range devices {
		switch strings.ToLower(device.Status) {
		case "completed":
			completed++
		case "failed":
			failed++
		case "pending":
			pending++
		}
	}
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Status", "# Devices"})
	t.AppendRow([]interface{}{"Completed", completed})
	t.AppendRow([]interface{}{"Failed", failed})
	t.AppendRow([]interface{}{"Pending", pending})
	util.RenderTable(t, format, false)

	// print devices impacted by the job
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Device ID", "Status"})
	var deviceNum int = 1
	for _, device := range devices {
		switch strings.ToLower(device.Status) {
		case "completed":
			if showCompleted {
				t.AppendRow([]interface{}{deviceNum, device.ID, device.Status})
				deviceNum++
			}
		case "failed":
			if showFailed {
				t.AppendRow([]interface{}{deviceNum, device.ID, device.Status})
				deviceNum++
			}
		case "pending":
			if showPending {
				t.AppendRow([]interface{}{deviceNum, device.ID, device.Status})
				deviceNum++
			}
		}
	}

	util.RenderTable(t, format, false)
}
