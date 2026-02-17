package main

import (
	"aggreGATOR/internal/database"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the follow command expects a single argument after the command, the url")
	}
	ctx := context.Background()

	// Get the feed it exists
	feed, err := s.db.GetFeedByURL(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}

	// Get current user
	currentUser, err := s.db.GetUser(ctx, s.config.CurrentUserName)
	if err != nil {
		return err
	}

	// Create feed follow with user and feed
	feedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UserID: currentUser.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("the following command does not need a url")
	}
	ctx := context.Background()
	currentUser, err := s.db.GetUser(ctx, s.config.CurrentUserName)
	if err != nil {
		return err
	}
	follows, err := s.db.GetFeedFollowsByUser(ctx, currentUser.ID)
	if err != nil {
		return err
	}
	for _, follow := range follows {
		fmt.Println(follow)
	}
	return nil

}
