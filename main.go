package main

import (
	"context"
	ps_utils "data_faker/pubsub_utils"
	"flag"
	"fmt"
)

var (
	project = flag.String("project", "", "Google Cloud Project")
	topic   = flag.String("topic", "", "Topic you wish to post to")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	client := ps_utils.CreateClient(ctx, *project)
	topic := client.Topic(*topic)

	if exists, _ := topic.Exists(ctx); !exists {
		fmt.Println("Fuck I dont exist")
	}

}
