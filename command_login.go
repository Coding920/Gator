package main

import (
	"context"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Must provide username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Not a username in the database")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User set")
	return nil
}
