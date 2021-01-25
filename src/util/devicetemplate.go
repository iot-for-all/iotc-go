package util

import (
	"com.azure.iot/iotcentral/iotcgo/models"
	"strings"
)

type ComponentTelemetry struct {
	Component string
	Telemetry []string
}

type ComponentCommands struct {
	Component string
	Commands  []string
}

// GetComponentTelemetry gets all the telemetry defined by all components including the default component
func GetComponentTelemetry(template *models.DeviceTemplate) []ComponentTelemetry {
	var components []ComponentTelemetry
	var defaultComponent ComponentTelemetry
	defaultComponent.Component = ""

	var implCapModels []interface{}
	cap, ok := template.CapabilityModel.(map[string]interface{})
	if ok {
		impl, ok := cap["implements"].([]interface{})
		if ok {
			contents, ok := cap["contents"].([]interface{})
			if ok {
				implCapModels = impl
				implCapModels = append(implCapModels, contents...)
			}
			for _, val := range implCapModels {
				comp, ok := val.(map[string]interface{})
				if ok {
					if comp["@type"] == "Component" {
						schema, ok := comp["schema"].(map[string]interface{})
						if ok {
							var ct ComponentTelemetry
							ct.Component = comp["name"].(string)
							contents := schema["contents"].([]interface{})
							for _, cont := range contents {
								//fmt.Printf("%T %v\n", cont, cont)
								con, ok := cont.(map[string]interface{})
								if ok {
									var typ, name string
									for contName, contVal := range con {
										if strings.ToLower(contName) == "@type" {
											typ, ok = contVal.(string)
											if !ok {
												types := contVal.([]interface{})
												for _, t := range types {
													tempType := strings.ToLower(t.(string))
													if tempType == "telemetry" {
														typ = tempType
													}
												}
											}
										} else if strings.ToLower(contName) == "name" {
											name = contVal.(string)
										}
									}
									if strings.ToLower(typ) == "telemetry" {
										ct.Telemetry = append(ct.Telemetry, name)
									}
								}
							}
							components = append(components, ct)
						}
					} else {
						var typ string
						types, ok := comp["@type"].([]interface{})
						if ok {
							for _, t := range types {
								tempType := strings.ToLower(t.(string))
								if tempType == "telemetry" {
									typ = tempType
								}
							}
							if typ == "telemetry" {
								defaultComponent.Telemetry = append(defaultComponent.Telemetry, comp["name"].(string))
							}
						}
					}
				}
			}
		}
	}

	// add default component telemetry
	if len(defaultComponent.Telemetry) > 0 {
		components = append(components, defaultComponent)
	}

	return components
}

// GetComponentCommands gets all the commands defined by all components including the default component
func GetComponentCommands(template *models.DeviceTemplate) []ComponentCommands {
	var commands []ComponentCommands
	var defaultComponent ComponentCommands
	defaultComponent.Component = ""

	var implCapModels []interface{}
	cap, ok := template.CapabilityModel.(map[string]interface{})
	if ok {
		impl, ok := cap["implements"].([]interface{})
		if ok {
			contents, ok := cap["contents"].([]interface{})
			if ok {
				implCapModels = impl
				implCapModels = append(implCapModels, contents...)
			}
			for _, val := range implCapModels {
				comp, ok := val.(map[string]interface{})
				if ok {
					if comp["@type"] == "Component" {
						schema, ok := comp["schema"].(map[string]interface{})
						if ok {
							var cc ComponentCommands
							cc.Component = comp["name"].(string)
							contents := schema["contents"].([]interface{})
							for _, cont := range contents {
								//fmt.Printf("%T %v\n", cont, cont)
								con, ok := cont.(map[string]interface{})
								if ok {
									var typ, name string
									for contName, contVal := range con {
										if strings.ToLower(contName) == "@type" {
											typ, ok = contVal.(string)
											if !ok {
												types := contVal.([]interface{})
												for _, t := range types {
													tempType := strings.ToLower(t.(string))
													if tempType == "command" {
														typ = tempType
													}
												}
											}
										} else if strings.ToLower(contName) == "name" {
											name = contVal.(string)
										}
									}
									if strings.ToLower(typ) == "command" {
										cc.Commands = append(cc.Commands, name)
									}
								}
							}
							commands = append(commands, cc)
						}
					} else if comp["@type"] == "Command" {
						defaultComponent.Commands = append(defaultComponent.Commands, comp["name"].(string))
					}
				}
			}
		}
	}

	// add default component commands
	if len(defaultComponent.Commands) > 0 {
		commands = append(commands, defaultComponent)
	}

	return commands
}
