package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rara-ch/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the register handler expects a single argument, the user's name")
	}

	name := cmd.args[0]

	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		os.Exit(1)
	}

	s.cfg.SetUser(name)

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		return fmt.Errorf("could not create user in db: %v", err)
	}

	fmt.Printf("The user was created: %v", user)

	return nil
}
