package devices

import (
	"com.azure.iot/iotcentral/iotcgo/client"
	"com.azure.iot/iotcentral/iotcgo/client/operations"
	"com.azure.iot/iotcentral/iotcgo/cmd/devicetemplates"
	"com.azure.iot/iotcentral/iotcgo/config"
	"com.azure.iot/iotcentral/iotcgo/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// getComponentsCmd represents the devices get last known value command
var getLKVCmd = &cobra.Command{
	Use:   "lkv",
	Short: "Get the last known values of telemetry, properties and cloud properties of the device",
	Long:  `Get the last known values of telemetry, properties and cloud properties of the device.`,
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
		deviceID, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		showTelemetry, err := cmd.Flags().GetBool("telemetry")
		if err != nil {
			return err
		}
		showProperties, err := cmd.Flags().GetBool("properties")
		if err != nil {
			return err
		}
		showCloudProperties, err := cmd.Flags().GetBool("cloudProperties")
		if err != nil {
			return err
		}

		if !showTelemetry && !showProperties && !showCloudProperties {
			return errors.New("at least one of these should be set to true - telemetry, properties, CloudProperties")
		}

		// create an IoTC API Client to connect to the given app
		c, err := client.NewFromToken(app)
		if err != nil {
			return err
		}

		// start the spinner
		spin := util.NewSpinner(" Downloading device ...")

		// get the device details
		deviceParams := operations.NewDevicesGetParams()
		deviceParams.DeviceID = deviceID
		res, err := c.Operations.DevicesGet(deviceParams)
		if err != nil {
			return err
		}

		if res.Payload == nil {
			fmt.Printf("No devices found in '%s' app\n", app)
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Type", "Value", "Timestamp"})

		// get all device templates look up table so that we can print the template names
		if showTelemetry {
			spin.Suffix = " Getting device template"
			deviceTemplate, err := devicetemplates.GetDeviceTemplate(c, app, res.Payload.InstanceOf)
			if err != nil || deviceTemplate == nil {
				return nil
			}

			// get telemetry values
			components := util.GetComponentTelemetry(deviceTemplate)
			for _, comp := range components {
				for _, telemetryName := range comp.Telemetry {
					if len(comp.Component) > 0 {
						ctvParams := operations.NewDevicesGetComponentTelemetryValueParams()
						ctvParams.DeviceID = deviceID
						ctvParams.ComponentName = comp.Component
						ctvParams.TelemetryName = telemetryName
						spin.Suffix = " Getting telemetry - " + comp.Component + "/" + telemetryName
						telemetry, err := c.Operations.DevicesGetComponentTelemetryValue(ctvParams)
						if err != nil {
							return err
						}
						b, _ := json.MarshalIndent(telemetry.Payload.Value, "", "  ")
						t.AppendRow(table.Row{comp.Component + "/" + telemetryName, string(b), telemetry.Payload.Timestamp})
					} else {
						tvParams := operations.NewDevicesGetTelemetryValueParams()
						tvParams.DeviceID = deviceID
						tvParams.TelemetryName = telemetryName
						spin.Suffix = " Getting telemetry - " + telemetryName
						telemetry, err := c.Operations.DevicesGetTelemetryValue(tvParams)
						if err != nil {
							return err
						}
						b, _ := json.MarshalIndent(telemetry.Payload.Value, "", "  ")
						t.AppendRow(table.Row{telemetryName, string(b), telemetry.Payload.Timestamp})
					}
				}
			}
		}

		// Get Properties
		if showProperties {
			spin.Suffix = " Getting device properties"
			propsParams := operations.NewDevicesGetPropertiesParams()
			propsParams.DeviceID = deviceID
			props, err := c.Operations.DevicesGetProperties(propsParams)
			if err != nil {
				return err
			}
			propsBytes, _ := json.MarshalIndent(props.Payload, "", "  ")
			t.AppendRow(table.Row{"Properties", string(propsBytes), ""})
		}

		// Get Cloud Properties
		if showCloudProperties {
			spin.Suffix = " Getting device cloud properties"
			cloudPropsParams := operations.NewDevicesGetCloudPropertiesParams()
			cloudPropsParams.DeviceID = deviceID
			cloudProps, err := c.Operations.DevicesGetCloudProperties(cloudPropsParams)
			if err != nil {
				return err
			}
			b, _ := json.MarshalIndent(cloudProps.Payload, "", "  ")
			t.AppendRow(table.Row{"Cloud Properties", string(b), ""})
		}

		spin.Stop()

		// write out the table
		util.RenderTable(t, format, false)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getLKVCmd)

	getLKVCmd.Flags().StringP("app", "a", "", "name of the IoT Central application")
	getLKVCmd.MarkFlagRequired("app")
	getLKVCmd.Flags().StringP("id", "", "", "unique device ID")
	getLKVCmd.MarkFlagRequired("id")
	getLKVCmd.Flags().StringP("format", "f", config.Config.Format, "output formats: pretty, table, csv, markdown, html")
	getLKVCmd.Flags().BoolP("telemetry", "", true, "get the telemetry last known values")
	getLKVCmd.Flags().BoolP("properties", "", true, "get the properties last known values")
	getLKVCmd.Flags().BoolP("cloudProperties", "", true, "get the cloud properties last known values")
}
