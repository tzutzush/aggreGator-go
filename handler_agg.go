package main

import (
	"aggreGATOR/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("aggregate command needs a single argument, the time between requests")
	}
	interval, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	timer := time.NewTicker(interval)
	defer timer.Stop()
	for ; ; <-timer.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}

	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("no feed found")
	}
	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("no rssfeed fetched")
	}
	fmt.Println("=======================================================================================================\n")
	for _, f := range rssFeed.Channel.Item {
		pubDate, parseErr := time.Parse(time.RFC1123Z, f.PubDate)
		if parseErr != nil {

		}
		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       f.Title,
			Url:         f.Link,
			Description: f.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			fmt.Println("Post already exists")
			continue
		} else {
			fmt.Println("New post saved")
		}
	}
	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("mark feed as fetched failed")
	}
	return nil
}
