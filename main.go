package main

import (
	"sphire/configuration"
	"fmt"
)

func main() {
	vpx, err := configuration.Configuration("DEV")
	if err != nil{
		panic(fmt.Errorf("Error reading configuration file: %s", err))
	}
	fmt.Println(vpx.GetString("application.log.error"))
}
