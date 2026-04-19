package main

import (
	"github.com/weilok2021/blog-aggregator/internal/config"
	"github.com/weilok2021/blog-aggregator/internal/database"
	"errors"
	"fmt"
	"time"
	"context"
	"os"
	"github.com/google/uuid"
	"net/http"
	"io"
	"encoding/xml"
	"html"
)

type state struct {
	db  *database.Queries
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

// To handle user login, if user already existed in db, exit with status 1
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The login handler expects a single argument, the username.")
	}
	
	ctx := context.Background()
	// Exit with code 1 if a user not exist in the database.
	if _, err := s.db.GetUser(ctx, cmd.args[0]); err != nil {
		fmt.Println("Couldn't login, no user record in database!")
		os.Exit(1)
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("Username has been set!")
	return nil
}

// To register a new user in users table
func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The register handler expects a single argument, the username.")
	}

	ctx := context.Background()
	// Exit with code 1 if a user with that name already exists.
	if _, err := s.db.GetUser(ctx, cmd.args[0]); err == nil {
		fmt.Println("User with this name already existed in database!")
		os.Exit(1)
	}

	createdUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	})
	if err != nil {
		return err
	}

	s.config.SetUser(cmd.args[0])
	fmt.Println("User has created! This is your information.")
	fmt.Println(createdUser)
	return nil
}

// To delete all users from users table
func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("The resit command expects no argument.")
	}
	ctx := context.Background()
	if err := s.db.Reset(ctx); err != nil {
		return err
	}
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("The users command expects no argument.")
	}
	
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		fmt.Println("Users table not existed in db!")
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("(current)")
		}
		fmt.Printf("\n")
	} 
	return nil
}

func handlerFeed(s *state, cmd command) error {
	ctx := context.Background()
	rssFeed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req.Header.Set("User-Agent", "Gator")

	res, err := client.Do(req) 
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body) 
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed 
	if err := xml.Unmarshal(bytes, &rssFeed); err != nil {
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
	return &rssFeed, nil
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

