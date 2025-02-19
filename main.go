package main

import (
	"internal/config"
	"fmt"
)

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Printf("Error in reading configuration: %s\n", err)
	}

	configErr := configuration.SetUser("megarage9000")
	if configErr != nil {
		fmt.Printf("Error in setting configuration: %s\n", configErr)
	}

	test, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	test.PrintConfig()
}