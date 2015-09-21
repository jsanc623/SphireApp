package configuration

import (
	//"github.com/spf13/viper"
    "strings"
	"fmt"
)

func Configuration(environment string) string {
	var cfx_file string = "development.json"

	switch environment {
	case "DEV":
		cfx_file = "development.json"
	case "STG":
		cfx_file = "staging.json"
	case "PROD":
		cfx_file = "production.json"
	}

	fmt.Println(strings.Split(cfx_file, ".")[0])

	//viper.SetConfigName(strings.Split(cfx_file, ".")[0]) // name of config file (without extension)
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	//viper.ReadInConfig() // Find and read the config file

	return ""
}
