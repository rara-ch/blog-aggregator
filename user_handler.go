package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}

	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		os.Exit(1)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return err
	}

	fmt.Println("The user has been set successfully")
	return nil
}

func handlerDisplayUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf("%s (current)", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
