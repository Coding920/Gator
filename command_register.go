package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Coding920/Gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Must provide name")
	}

	currentTime := time.Now()

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User Created")
	fmt.Printf("%+v\n", user)

	return nil
}
