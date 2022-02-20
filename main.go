package main

import (
	"context"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/kamilwoloszyn/sample-gql-api/config"
	"github.com/kamilwoloszyn/sample-gql-api/infrastucture/storage/db"
	"github.com/kamilwoloszyn/sample-gql-api/infrastucture/storage/db/repo"
)

func main() {
	ctx := context.Background()
	conf := config.Config{}
	if err := env.Parse(&conf); err != nil {
		log.Fatalf("unable to parse config: %s", err)
	}

	database, err := db.InitializeDatabase(
		ctx,
		conf.DatabaseConnectionTimeout,
		conf.DatabaseHost,
		conf.DatabasePort,
		conf.DatabaseName,
	)
	if err != nil {
		log.Fatalf("initialize db: %s", err)
	}

	if err := database.TestConnection(ctx, conf.DatabaseConnectionTimeout); err != nil {
		log.Fatalf("ping database failed: %s", err)
	}

	categoryRepo := repo.NewCategoryRepo(database)
	productRepo := repo.NewProductRepo(database)

	database.SetRepo(
		db.SetCategoryRepo(categoryRepo),
		db.SetProductRepo(productRepo),
	)

}
