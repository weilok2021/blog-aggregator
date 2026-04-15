package main

import (
	"fmt"
	"github.com/weilok2021/blog-aggregator/internal/config"
)

func main() {
	cfg, _ := config.Read() 
	s := state{config: &cfg}
	c := command{name: "login", args: []string{"Wei Lok"}}

	handlerLogin(&s, c)
	fmt.Println(cfg)
}