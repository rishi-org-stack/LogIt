package github

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func cloudStorage() {
	// [START cloud_storage_golang]
	config := &firebase.Config{
		StorageBucket: "<BUCKET_NAME>.appspot.com",
	}
	opt := option.WithAPIKey(apiKey)
	// opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	log.Printf("Created bucket handle: %v\n", bucket)
}

const (
	apiKey = "AIzaSyCmZ7GjjbntGqkqpRs0vd9A3P298-jCf5Y"

	authDomain = "logit-52911.firebaseapp.com"

	projectId = "logit-52911"

	storageBucket = "logit-52911.appspot.com"

	messagingSenderId = "1027016320708"

	appId = "1:1027016320708:web:b4c7abcc6dc60931130b91"

	measurementId = "G-0WY88TQRLP"
)
