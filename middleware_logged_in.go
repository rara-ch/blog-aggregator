package main

import (
	"context"
	"fmt"

	"github.com/rara-ch/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("could not get user from database: %v", err)
		}

		if err := handler(s, cmd, user); err != nil {
			return fmt.Errorf("could not execute handler: %v", err)
		}

		return nil
	}
}
