package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlerFuncs map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handlerFunc, ok := c.handlerFuncs[cmd.name]; !ok {
		return errors.New("command does not exist")
	} else {
		return handlerFunc(s, cmd)
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFuncs[name] = f
}
