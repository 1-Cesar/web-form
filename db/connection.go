package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func getConnection() (client *mongo.Client, ctx context.Context) {
	var cred options.Credential

	cred.Username = "root"
	cred.Password = "root"

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return
}
