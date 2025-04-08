package main

import (
	"internal/config"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
	"fmt"
	"os"
	"database/sql"
)

import _ "github.com/lib/pq"

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Printf("Error in reading configuration: %s\n", err)
	}
	// Creating database connection
	db, err := sql.Open("postgres", configuration.DBUrl)
	dbQueries := database.New(db)

	// We can create pointers using the & operator
	current_state := &state {
		config: &configuration,
		database: dbQueries,
	}

	commands_list := commands {
		command_map: map[string]func(*state, command) error {
			"login": handlerLogin,
			"register": handlerRegister,
			"reset": handlerReset,
			"users": handlerUsers,
			"agg": handlerAgg,
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