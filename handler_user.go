package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("missing username")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Printf("User set to %s\n", cmd.Args[0])
	return nil
}
