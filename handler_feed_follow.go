package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ewielguszewski/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("missing URL: usage `follow <url>`")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed does not exist")
		}
		return fmt.Errorf("error checking feed: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	fmt.Printf("Feed follow created successfully:\n")

	fmt.Println("* Feed: ", follow.FeedName)
	fmt.Println("* Username: ", follow.UserName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching feed follows: %w", err)
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("* Feed: %s\n", follow.FeedName)
	}

	return nil
}
