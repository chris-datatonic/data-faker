package pubsub_utils

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func CreateClient(ctx context.Context, project string) *pubsub.Client {
	client, err := pubsub.NewClient(ctx, project)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
