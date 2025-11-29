package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get feeds from database: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %v, URL: %v, User: %v", feed.FeedName, feed.Url, feed.UserName)
	}

	return nil
}
