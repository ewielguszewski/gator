package main

import (
	"context"
	"fmt"

	"github.com/ewielguszewski/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("user not set")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
