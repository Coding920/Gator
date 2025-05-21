package main

import (
	"context"
	"fmt"
)

func resetUsersDb(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Experienced problem deleting users: %v", err)
	}
	fmt.Println("DB Reset")
	return nil
}
