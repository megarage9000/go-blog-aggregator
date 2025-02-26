package main

import (
	"internal/config"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
)

type state struct {
	config *config.Config
	database *database.Queries
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	// A map mapping strings to a funct(*state, command) error function
	command_map map[string]func(*state, command) error
}