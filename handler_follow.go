package main

import (
	"aggreGATOR/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the follow command expects a single argument after the command, the url")
	}
	ctx := context.Background()

	// Get the feed it exists
	feed, err := s.db.GetFeedByURL(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}

	// Create feed follow with user and feed
	feedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("the following command does not need a url")
	}
	ctx := context.Background()
	follows, err := s.db.GetFeedFollowsByUser(ctx, user.ID)
	if err != nil {
		return err
	}
	for _, follow := range follows {
		fmt.Println(follow)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the unfollow command needs a url")
	}
	ctx := context.Background()
	feed, err := s.db.GetFeedByURL(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollowFromUser(ctx, database.DeleteFeedFollowFromUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	return err
}
