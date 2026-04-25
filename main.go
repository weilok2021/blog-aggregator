package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"github.com/weilok2021/blog-aggregator/internal/config"
	"os"
	"database/sql"
	"github.com/weilok2021/blog-aggregator/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("At least 2 CLI arguments are needed!")
		os.Exit(1)
	}

	cfg, _ := config.Read() 
	appState := &state{config: &cfg}

	// Create database connection
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	appState.db = database.New(db) // we use queries through state.db

	cmds := newCommands()                                                                                             
	cmds.register("login", handlerLogin)             
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("addfeed", middlewareLoggedIn(addfeed))
	cmds.register("feeds", listfeeds)
	cmds.register("follow", middlewareLoggedIn(addFollow))
	cmds.register("following", middlewareLoggedIn(listUserFollowing))
	cmds.register("unfollow", middlewareLoggedIn(unfollow))
	cmds.register("agg", handlerAgg)
	cmds.register("browse", middlewareLoggedIn(browse))


	cmd := command{name: os.Args[1], args: os.Args[2:]}
	if err := cmds.run(appState, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}