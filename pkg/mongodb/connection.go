package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/diazharizky/go-graphql-bootstrap/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	config.Global.SetDefault("mongodb.host", "localhost")
	config.Global.SetDefault("mongodb.port", 27017)
	config.Global.SetDefault("mongodb.db", "gqlbootstrap")
}

func GetDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf(
		"mongodb://%s:%d",
		config.Global.GetString("mongodb.host"),
		config.Global.GetInt("mongodb.port"),
	)

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	db := config.Global.GetString("mongodb.db")

	return client.Database(db), nil
}
