package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commandRegistry struct {
	commands map[string]func(*state, command) error
}

func (c *commandRegistry) run(s *state, cmd command) error {
	commandToRun, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("command with name %s does not exist", cmd.name)
	}
	return commandToRun(s, cmd)
}

func (c *commandRegistry) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
