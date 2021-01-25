package devicetemplates

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/models"
	"com.azure.iot/iotcentral/iotcgo/util"
	"fmt"
	"github.com/jedib0t/go-pretty/progress"
	"github.com/spf13/cobra"
	"time"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy all the device templates from an application to another",
	Long:  `Copy all the device templates from an application to another.
You can either overwrite existing templates or skip if they exist.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// read the command line parameters
		srcApp, err := cmd.Flags().GetString("srcApp")
		if err != nil {
			return err
		}
		destApp, err := cmd.Flags().GetString("destApp")
		if err != nil {
			return err
		}
		overwrite, err := cmd.Flags().GetBool("overwrite")
		if err != nil {
			return err
		}

		// instantiate a Progress Writer and set up the options
		pw := progress.NewWriter()
		pw.SetAutoStop(false)
		pw.SetTrackerLength(20)
		pw.ShowOverallTracker(false)
		pw.ShowTime(true)
		pw.ShowTracker(true)
		pw.ShowValue(false)
		pw.SetMessageWidth(50)
		pw.SetNumTrackersExpected(2)
		pw.SetSortBy(progress.SortByPercentDsc)
		pw.SetStyle(progress.StyleDefault)
		pw.SetTrackerPosition(progress.PositionRight)
		pw.SetUpdateFrequency(time.Millisecond * 100)
		pw.Style().Colors = progress.StyleColorsExample
		pw.Style().Options.PercentFormat = "%4.1f%%"

		// call Render() in async mode; yes we don't have any trackers at the moment
		go pw.Render()

		// get device templates from source app
		srcTemplates, err := getDeviceTemplates(srcApp, pw)
		if err != nil {
			return err
		}

		// get device templates from destination app
		destTemplates, err := getDeviceTemplates(destApp, pw)
		if err != nil {
			return err
		}

		// progress trackers for each template
		var dtTracker = make([]progress.Tracker, len(srcTemplates))
		for i := 0; i < len(srcTemplates); i++ {
			dtTracker[i] = progress.Tracker{
				Message: srcTemplates[i].DisplayName,
				Total:   100,
				Units:   progress.UnitsDefault}
		}

		////////////////////////////////////////////////////////////////////////////////
		// Upload new device templates from source app to destination app
		////////////////////////////////////////////////////////////////////////////////
		destClient, err := client.NewFromToken(destApp)
		if err != nil {
			return err
		}
		for i := 0; i < len(srcTemplates); i++ {
			pw.AppendTracker(&dtTracker[i])

			var found bool = false
			if destTemplates != nil {
				for _, destDT := range destTemplates {
					if destDT.ID == srcTemplates[i].ID {
						found = true
						break
					}
				}
			}

			// Add a new template if it does not exist in destination application
			// Update the template if it exists
			newDTParams := operations.NewDeviceTemplatesSetParams()
			newDTParams.DeviceTemplateID = srcTemplates[i].ID
			newDTParams.Body = srcTemplates[i]
			newDTParams.Body.Etag = ""
			mode := "Updated"
			if !found {
				newDTParams.Body.ID = ""
				mode = "Added"
			}

			if found && !overwrite {
				mode = "Skipped"
			} else {
				_, err := destClient.Operations.DeviceTemplatesSet(newDTParams)
				if err != nil {
					return err
				}
			}
			dtTracker[i].Message = fmt.Sprintf("%-30s %10s", srcTemplates[i].DisplayName, mode)
			dtTracker[i].Increment(100)
			dtTracker[i].MarkAsDone()
		}

		time.Sleep(time.Millisecond * 1000)

		return nil
	},
}

func init() {
	deviceTemplatesCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringP("srcApp", "s", "", "source application to copy from")
	copyCmd.MarkFlagRequired("srcApp")
	copyCmd.Flags().StringP("destApp", "d", "", "destination application to copy into")
	copyCmd.MarkFlagRequired("destApp")
	copyCmd.Flags().BoolP("overwrite", "", false, "skip copy if the template exists i.e. do not overwrite it")
}

func getDeviceTemplates(app string, pw progress.Writer) ([]*models.DeviceTemplate, error){
	var dt []*models.DeviceTemplate
	client, err := client.NewFromToken(app)
	if err != nil {
		return nil, err
	}

	tracker := progress.Tracker{Message: fmt.Sprintf("Downloading from %s", app), Total: 100, Units: progress.UnitsDefault}
	pw.AppendTracker(&tracker)

	templates, err := client.Operations.DeviceTemplatesList(operations.NewDeviceTemplatesListParams())
	if err != nil {
		return nil, err
	}
	dt = append(dt, templates.Payload.Value...)

	// loop through and download all the rows one page at a time
	nextLink := templates.Payload.NextLink
	for {
		if len(nextLink) == 0 {
			break
		}

		numDT := len(dt)
		tracker.Message = fmt.Sprintf("Downloading from %s (%v)", app, numDT)
		body, err := util.GetContent(app, nextLink)
		if err != nil {
			return nil, err
		}

		var dc models.DeviceTemplateCollection
		if err := dc.UnmarshalBinary(body); err != nil {
			return nil, err
		}
		dt = append(dt, dc.Value...)

		nextLink = dc.NextLink
	}
	tracker.Increment(100)
	tracker.Message = fmt.Sprintf("Downloaded %v from %-22s", len(dt), app)

	return dt, nil
}

