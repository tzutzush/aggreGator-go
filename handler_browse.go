package main

import (
	"aggreGATOR/internal/database"
	"context"
	"fmt"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("browse command can have 1 optional argument, a limit, base limit is 2")
	}

	ctx := context.Background()
	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  2,
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
	}
	return nil
}
