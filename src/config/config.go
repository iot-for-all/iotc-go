package config

import "errors"

// YamlConfig is exported.
type App struct {
	Name      string `yaml:"name"`
	Subdomain string `yaml:"subdomain"`
	Domain    string `yaml:"domain"`
	ApiToken  string `yaml:"apiToken"`
}
type IoTCConfig struct {
	Apps    []App	`yaml:"apps"`
	MaxRows int		`yaml:"maxRows"`
	Format  string	`yaml:"format"`
}

var (
	Config IoTCConfig
)

func init() {
	if Config.MaxRows == 0 {
		Config.MaxRows = 25
	}
	if Config.Format == "" {
		Config.Format = "pretty"
	}
}

func GetAppConfig(appName string) (App, error) {
	for _, value := range Config.Apps {
		if value.Name == appName {
			return value, nil
		}
	}

	return App{}, errors.New("Could not find app '" + appName + "' in configuration file")
}