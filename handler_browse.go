package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ewielguszewski/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if providedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = providedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	fmt.Printf("Posts for user %s:\n", user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("2006-01-02 15:04:05"), post.FeedName)
		fmt.Printf("----- %s -----\n", post.Title)
		fmt.Printf("   %s\n", post.Description.String)
		fmt.Printf("   %s\n", post.Url)
	}
	return nil
}
