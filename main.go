package main

import (
	"flag"
	"log"
)

var (
	configPath = flag.String("config", "./config.json", "Path to configuration JSON")
)

func main() {
	flag.Parse()

	log.Println(*configPath)
}
