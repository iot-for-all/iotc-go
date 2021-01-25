package cmd

import (
	"com.azure.iot/iotcentral/iotcgo/config"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var appConfigFile string

// RootCmd is the root command.
var RootCmd = &cobra.Command{
	Use: "iotc",
	Short: "Azure IoT Central CLI (GO)",
	Long: `
    ____    ______   ______           __             __
   /  _/___/_  __/  / ____/__  ____  / /__________  / /
   / // __ \/ /    / /   / _ \/ __ \/ __/ ___/ __ \/ / 
 _/ // /_/ / /    / /___/  __/ / / / /_/ /  / /_/ / /  
/___/\____/_/     \____/\___/_/ /_/\__/_/   \__,_/_/   
                                                       
This tool helps you manage IoT Central applications through simple CLI 
commands. You can explore (list, upload, download) several entities such
as devicesTemplates, roles, users in an IoT Central applications.
You can copy these entities between different apps. You an operate against
each of these entities below using sub-commands underneath each command.`,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	initConfig()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Command config file (default is $HOME/.iotc.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".iotc.yml")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%v\n", err)
		var format = `Add a configuration file $HOME/.iotc.yml with the file format:

apps:
  - name: myapp1
    subdomain: app1
    appToken: app1Token
  - name: myapp2
    subdomain: app2
    appToken: app2Token

# Maximum number of rows to retrieve for 'list' commands
maxRows: 25

# Default format of the tables
format: pretty
`
		fmt.Print(format)
	} else {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	initAppConfig()
}


func initAppConfig() {
	/*configFile, err := ioutil.ReadFile(appConfigFile)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
	}

	////////////////////////////////////////////////////////////////////////////////
	// parse configuration file
	////////////////////////////////////////////////////////////////////////////////
	err = yaml.Unmarshal(configFile, &config.Config)
	if err != nil {
		fmt.Printf("Error parsing app config file: %s\n", err)
		os.Exit(1)
	}*/

	err := viper.Unmarshal(&config.Config)
	if err != nil {
		fmt.Printf("Error parsing app config file: %s\n", err)
		os.Exit(1)
	}
}
