package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/kamilwoloszyn/sample-gql-api/config"
)

func main() {
	conf := config.Config{}
	if err := env.Parse(&conf); err != nil {
		log.Fatalf("unable to parse config: %s", err)
	}
}
