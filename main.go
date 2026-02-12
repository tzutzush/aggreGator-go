package main

import (
	"aggreGATOR/internal/config"
	"fmt"
	"os"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("reading config unsuccessful")
	}
	applicationState := &state{config: &cfg}
	cmds := commandRegistry{commands: make(map[string]func(*state, command) error)}
	cmds.register("login", loginHandler)

	if len(os.Args) < 2 {
		fmt.Println("no argument was provided")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(applicationState, command{name: cmdName, arguments: cmdArgs})
	if err != nil {
		fmt.Println("failed to run command")
		os.Exit(1)
	}
}
