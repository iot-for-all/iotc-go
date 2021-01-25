package main

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/actions"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/apitokens"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/apps"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/cde"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/completion"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/devicegroups"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/devices"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/devicetemplates"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/jobs"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/roles"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/rules"
	_ "com.azure.iot/iotcentral/iotcgo/cmd/users"
)

func main() {
	cmd.Execute()
}