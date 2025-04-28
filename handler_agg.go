package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ewielguszewski/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reps>", cmd.Name)
	}

	timeBetweenReps, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s\n", timeBetweenReps)

	ticker := time.NewTicker(timeBetweenReps)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("error fetching next feed: %w", err)
		return
	}
	log.Println("Found next feed to fetch")

	scrapeFeed(s, feed)
}

func scrapeFeed(s *state, feed database.Feed) {

	_, err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %w", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("error fetching feed: %w", err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found.", feed.Name, len(feedData.Channel.Item))
}
