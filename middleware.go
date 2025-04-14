package main

import (
	"context"
	"github.com/megarage9000/go-blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s * state, cmd command, user database.User) error) func (*state, command) error {
	return func(s * state, cmd command) error {
		
		user, err := s.database.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}
		
		// Call the wrapped function
		return handler(s, cmd, user)
	}
}