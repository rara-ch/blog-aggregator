package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("could not get user from database: %v", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get follows from database: %v", err)
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
