package main

import (
	"basic-api/internal/app"
	db "basic-api/internal/core/database"
	"basic-api/internal/core/logger"
	"basic-api/internal/core/server"
	"basic-api/internal/env"
	"basic-api/internal/router"

	"github.com/sirupsen/logrus"
)

func main() {
	envConfig, err := env.Read("configs")
	if err != nil {
		panic(err)
	}

	ds, err := db.NewMariaDB(&db.Config{
		Host:         envConfig.Config.Database.Host,
		Port:         envConfig.Config.Database.Port,
		User:         envConfig.Config.Database.Username,
		Password:     envConfig.Config.Database.Password,
		DatabaseName: envConfig.Config.Database.Name,
		Debug:        !envConfig.Config.Release,
	})
	if err != nil {
		panic(err)
	}

	mds, err := db.NewMongoDB(&db.Config{
		Host:         envConfig.Config.Mongo.Host,
		User:         envConfig.Config.Mongo.Username,
		Password:     envConfig.Config.Mongo.Password,
		DatabaseName: envConfig.Config.Mongo.Name,
	})
	if err != nil {
		panic(err)
	}

	options := &router.Options{
		AppContext: &app.Context{
			Db:      ds.Db,
			Config:  envConfig.Config,
			MongoDb: mds,
		},
	}
	logrus.SetOutput(&logger.OutputSplitter{})

	server.New(router.NewWithOptions(options), envConfig.Config.ServerPort).Start()
}
