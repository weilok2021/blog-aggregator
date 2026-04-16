package main

import (
	"fmt"
	"github.com/weilok2021/blog-aggregator/internal/config"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("At least 2 CLI arguments are needed!")
		os.Exit(1)
	}

	cfg, _ := config.Read() 
	appState := &state{config: &cfg}
	cmds := newCommands()                                                                                             
	cmds.register("login", handlerLogin)             

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	err := cmds.run(appState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cfg)
	// start CLI application
	// for {
	// 	if len(os.Args) < 2 {
	// 		fmt.Errorf("At least 2 CLI arguments are needed!")
	// 		os.Exit(1)
	// 	}
	// }
}