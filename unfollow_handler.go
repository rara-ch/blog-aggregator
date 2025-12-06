package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/rara-ch/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("the unfollow handler expects a single argument, the url")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed from database: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not delete feed follow from database: %v", err)
	}

	fmt.Printf("User %v unfollowed %v", user.Name, feed.Name)

	return nil
}
