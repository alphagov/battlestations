package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/alphagov/battlestations/api"
	"github.com/alphagov/battlestations/github"
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
		os.Exit(1)
	}

	githubService := github.NewOAuthService(
		fmt.Sprintf("%s/authorized", config.BaseURL),
		config.Github,
	)

	authKey, _ := ioutil.ReadFile(config.AuthKey)
	encKey, _ := ioutil.ReadFile(config.EncKey)

	log.Printf("API starting at http://%s", config.Addr)
	http.ListenAndServe(
		config.Addr,
		api.MakeRouter(authKey, encKey, githubService))

}
