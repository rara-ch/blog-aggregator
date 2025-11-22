package main

import (
	"fmt"
	"log"

	"github.com/rara-ch/blog-aggregator/internal/config"
)

func main() {
	doConfig()
}

func doConfig() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	cfg.CurrentUsername = username
	err = cfg.SetUser()
	if err != nil {
		log.Fatalf("unable to set username: %v", err)
	}

	newCfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}
	fmt.Println(newCfg)
}
