package main

import (
	"aggreGATOR/internal/database"
	"context"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		currentUser, err := s.db.GetUser(ctx, s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, currentUser)
	}
}
