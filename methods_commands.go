package main

import (
	"fmt"
)

func handlerLogin(s * state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("error, no arguments provided")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to %s\n", cmd.arguments[0])
	return nil
}

// Command methods
func (c * commands) register(name string, f func(*state, command) error) {
	c.command_map[name] = f
}

func (c * commands) run(s * state, cmd command) error {
	if s == nil {
		return fmt.Errorf("error: no state exists")
	}
	err := c.command_map[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}