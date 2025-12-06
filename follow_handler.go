package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rara-ch/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the follow handler expects a single argument, the url")
	}

	url := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("could not get user from database: %v", err)
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed from database: %v", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow in db: %v", err)
	}

	fmt.Printf("Feed: %s\n", feed_follow.FeedName)
	fmt.Printf("Name: %s", feed_follow.UserName)

	return nil
}
