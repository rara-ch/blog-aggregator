package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("could not reset user table: %v", err)
	} else {
		fmt.Println("Reset user table")
		return nil
	}
}
