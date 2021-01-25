package jobs

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"github.com/spf13/cobra"
)

// stopCmd represents the jobs stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop execution of a job that is currently running",
	Long: `Stop execution of a job that is currently running`,
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

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// stop the job
		p := operations.NewJobsStopParams()
		p.JobID = id
		_, err = c.Operations.JobsStop(p)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	jobsCmd.AddCommand(stopCmd)

	stopCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	stopCmd.MarkFlagRequired("app")
	stopCmd.Flags().StringP("id", "", "", "unique ID for the job")
	stopCmd.MarkFlagRequired("id")
}
