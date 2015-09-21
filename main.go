package main

import (
	"sphire/configuration"
	"fmt"
)

func main() {
	cfx := configuration.Configuration("DEV")
	fmt.Println(cfx)
}
