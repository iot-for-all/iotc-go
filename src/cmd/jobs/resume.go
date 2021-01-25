package jobs

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"github.com/spf13/cobra"
)

// resumeCmd represents the jobs resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume execution of an existing stopped job",
	Long: `Resume execution of an existing stopped job`,
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

		// Resume the job
		p := operations.NewJobsResumeParams()
		p.JobID = id
		_, err = c.Operations.JobsResume(p)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	jobsCmd.AddCommand(resumeCmd)

	resumeCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	resumeCmd.MarkFlagRequired("app")
	resumeCmd.Flags().StringP("id", "", "", "unique ID for the job")
	resumeCmd.MarkFlagRequired("id")
}
