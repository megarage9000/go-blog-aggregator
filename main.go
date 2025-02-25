package main

import (
	"internal/config"
	"fmt"
	"os"
)

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Printf("Error in reading configuration: %s\n", err)
	}

	// We can create pointers using the & operator
	current_state := &state {
		config: &configuration,
	}

	commands_list := commands {
		command_map: map[string]func(*state, command) error {
			"login": handlerLogin,
		},
	}

	args := os.Args 

	if len(args) < 2 {
		fmt.Printf("error in retrieving arguments\n")
		os.Exit(1)
	}

	entered_command := command {
		name: args[1],
		arguments: args[2:],
	}

	if err := commands_list.run(current_state, entered_command); err != nil{
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}