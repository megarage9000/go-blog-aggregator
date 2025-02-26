package main

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
)

func handlerLogin(s * state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("error, no arguments provided")
	}

	name := cmd.arguments[0]
	if err := checkIfUserExists(s, name); err == nil {
		return fmt.Errorf("error, user %s does not exist\n", name)
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to %s\n", cmd.arguments[0])
	return nil
}

func handlerRegister(s * state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("error, no arguments provided")
	}

	name := cmd.arguments[0]
	if err := checkIfUserExists(s, name); err != nil {
		return err
	}

	// Create user
	userParams := database.CreateUserParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}
	user, err := s.database.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	// Set user
	config_err := s.config.SetUser(user.Name)
	if config_err != nil {
		return config_err
	}

	fmt.Printf("user has been for %s and has been set to %s\n", user.Name, user.Name)
	return nil
}

func checkIfUserExists(s * state, name string) error {
	// Check if user exists, if it does return error
	_, err := s.database.GetUser(context.Background(), name)

	if err == sql.ErrNoRows {
		// User does not exist, we want this!
		return nil
	}

	if err != nil {
		return fmt.Errorf("error in Getting User: %w\n", err)
	}

	// If we get here, the user was found (no errors and rows exist)
	// We should return an error because the user already exists
	return fmt.Errorf("user %s already exists", name)
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