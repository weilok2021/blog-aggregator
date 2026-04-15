package main

import (
	"fmt"
	"github.com/weilok2021/blog-aggregator/internal/config"
)

func main() {
	cfg, _ := config.Read() 
	cfg.SetUser("Lane")
	newConfig, _ := config.Read()
	fmt.Println(newConfig)
}