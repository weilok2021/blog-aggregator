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
	cmds.register("agg", handlerFeed)
	cmds.register("addfeed", addfeed)
	cmds.register("feeds", listfeeds)
	cmds.register("follow", addFollow)
	cmds.register("following", listUserFollowing)

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	if err := cmds.run(appState, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// start CLI application
	// for {
	// 	if len(os.Args) < 2 {
	// 		fmt.Errorf("At least 2 CLI arguments are needed!")
	// 		os.Exit(1)
	// 	}
	// }
}