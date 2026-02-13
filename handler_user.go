package main

import (
	"aggreGATOR/internal/database"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("user does not exist")
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("setting username failed")
	}
	fmt.Println("User has been set.")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the register handler expects a single argument after the command, the username")
	}
	name := cmd.arguments[0]
	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name: name,
	})
	if err != nil {
		return err
	}
	err = s.config.SetUser(name)
	if err != nil {
		return err
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("reset command does not need arguments")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Reset successful")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("users command does not need arguments")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
