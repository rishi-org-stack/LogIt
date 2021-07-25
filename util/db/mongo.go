package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"logit/v1/util/config"
)

func Connect(ctx context.Context, env *config.Env) *mongo.Client {
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(env.DB),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.", client)
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return client
}
