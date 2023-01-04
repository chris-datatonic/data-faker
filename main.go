package main

import (
	"data_faker/faker"
	"data_faker/utils"
	"flag"
	"fmt"
)

var (
	project    = flag.String("project", "", "Google Cloud Project")
	topic_name = flag.String("topic", "", "Topic you wish to post to")
	input      = flag.String("input", "", "Json file that contains faker structure")
	num_values = flag.Int("num_values", 1, "Number of values you want to produce")
)

func main() {
	flag.Parse()

	data := utils.MarshalFileToJson(*input)
	customStructure := faker.MakeStructure(data)

	for i := 0; i < *num_values; i++ {
		v := faker.CreateFakeData(customStructure)
		fmt.Println(v)
	}

	// ctx := context.Background()
	// client := ps_utils.CreateClient(ctx, *project)
	// topic := client.Topic(*topic_name)

	// exists, err := topic.Exists(ctx)

	// if err != nil {
	// 	log.Fatalf("Failed to check topic %v\n Error: %v.", *topic_name, err)
	// }

	// if !exists {
	// 	fmt.Println("Fuck I dont exist")
	// }

}
