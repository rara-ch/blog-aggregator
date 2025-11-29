package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) < 1 {
	// 	return errors.New("the agg handler expects a single argument, the url")
	// }

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(feed)
	return nil
}
