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
