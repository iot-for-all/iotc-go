package devicetemplates

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
)

// deviceTemplatesCmd represents the deviceTemplates command
var deviceTemplatesCmd = &cobra.Command{
	Use:   "deviceTemplates",
	Short: "Create, read, copy (between apps) and delete device templates within an IoT Central application",
	Long: `Device templates are composed of device capability model and the solution model.
Create, read, copy (between apps) and delete device templates within an IoT Central application.`,
}

func init() {
	cmd.RootCmd.AddCommand(deviceTemplatesCmd)
}
