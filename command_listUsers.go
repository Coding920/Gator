package main

import (
	"context"
	"fmt"
)

func listUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.cfg.CurrentUserName == user {
			fmt.Printf("%v (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
