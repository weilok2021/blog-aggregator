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

type commandHandler func(*state, command) error                                                                   
   
type commands struct {                                                                                            
	handlers map[string]commandHandler
}                                                                                                                 

func newCommands() *commands {                                                                                    
	return &commands{handlers: make(map[string]commandHandler)}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The login handler expects a single argument, the username.")
	}
	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("Username has been set!")
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	// This method runs a given command with the provided state if it exists.
	funcHandler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("Command not exist in this app!")
	}
	return funcHandler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	// This method registers a new handler function for a command name.
	c.handlers[name] = f
	return nil
}
