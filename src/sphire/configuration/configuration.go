package configuration

import (
	"github.com/spf13/viper"
	"strings"
)

//type vpx *viper.Viper

func Configuration(environment string) (*viper.Viper, error) {
	var cfx_file string = "development.json"

	switch environment {
	case "STG":
		cfx_file = "staging.json"
	case "PRD":
		cfx_file = "production.json"
	}

	var vpx *viper.Viper = viper.New()
	vpx.SetConfigName(strings.Split(cfx_file, ".")[0]) // name of config file (without extension)
	vpx.AddConfigPath("/opt/sphire/config")            // path to look for the config file in
	err := vpx.ReadInConfig()                          // Find and read the config file

	if err != nil { // Handle errors reading the config file
		return vpx, err
	}

	return vpx, nil
}
