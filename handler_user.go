package main

import "fmt"

func loginHandler(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("setting username failed")
	}
	fmt.Println("User has been set.")
	return nil
}
