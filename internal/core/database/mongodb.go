package database

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewMongoDB(conf *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dns := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		conf.User,
		conf.Password,
		conf.Host,
	)

	opts := options.Client()
	opts.ApplyURI(dns)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Errorf("cannot connect to mongodb err : %s", err)
		return nil, err
	}

	return client.Database(conf.DatabaseName), nil
}
