package jobs

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Operate against IoT Central jobs",
	Long: `Operate against IoT Central jobs.`,
}

func init() {
	cmd.RootCmd.AddCommand(jobsCmd)
}