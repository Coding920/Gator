package main

import (
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Must provide username")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User set")
	return nil
}
