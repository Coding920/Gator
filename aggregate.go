package main

import (
	"context"
	"fmt"
)

func agg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Please provide Url")
	}

	feed, err := fetchFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
