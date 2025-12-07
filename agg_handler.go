package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the agg handler expects a single argument, the time between requests")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("could not parse time between requests: %v", err)
	}

	fmt.Printf("Collecting feeds every %v", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return err
	// }

	// fmt.Print(feed)
	// return nil
}

func scrapeFeeds(s *state) error {
	feedPrev, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed to fetch from database: %v", err)
	}

	feedCurr, err := fetchFeed(context.Background(), feedPrev.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %v", err)
	}

	if _, err := s.db.MarkFeedAsFetched(context.Background(), feedPrev.ID); err != nil {
		return fmt.Errorf("could not mark feed as fetched: %v", err)
	}

	for _, item := range feedCurr.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
