package main

import (
	"fmt"
	"log"

	"github.com/rara-ch/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.CurrentUsername = username
	cfg.SetUser()
	newCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newCfg)
}
