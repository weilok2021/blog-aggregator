package main

import (
	"github.com/weilok2021/blog-aggregator/internal/config"
	"errors"
	"fmt"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("The login handler expects a single argument, the username.")
	}
	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("Username has been set!")
	return nil
}