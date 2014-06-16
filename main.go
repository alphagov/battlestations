package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/alphagov/battlestations/api"
)

var (
	configPath = flag.String("config", "./config.json", "Path to configuration JSON")
)

func main() {

	flag.Parse()

	err, config := ParseConfig(*configPath)

	if err != nil {
		log.Printf("Failed to read config from %s", *configPath)
		log.Println(err)
	} else {
		log.Printf("API starting at http://%s", config.Addr)
		http.ListenAndServe(config.Addr, api.MakeRouter())
	}

}
