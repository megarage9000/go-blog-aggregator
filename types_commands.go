package main

import (
	"internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	// A map mapping strings to a funct(*state, command) error function
	command_map map[string]func(*state, command) error
}