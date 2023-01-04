package faker

import (
	"data_faker/utils"
	"log"
	"reflect"
	"strings"

	"github.com/go-faker/faker/v4"
)

func getType(typeKey string) (t reflect.Type) {
	switch strings.ToLower(typeKey) {
	case "string":
		t = reflect.TypeOf("")
	case "int", "integer":
		t = reflect.TypeOf(0)
	case "bool", "boolean":
		t = reflect.TypeOf(true)
	default:
		log.Fatalf("type %v unsupported", typeKey)
	}
	return
}

func MakeStructure(json []utils.Config) (newStruct reflect.Type) {
	sfs := make([]reflect.StructField, 0)

	for _, config := range json {
		sf := reflect.StructField{
			Name: config.Name,
			Type: getType(config.JsonType),
			Tag:  reflect.StructTag(config.FakerType),
		}
		sfs = append(sfs, sf)
	}

	newStruct = reflect.StructOf(sfs)
	return
}

func CreateFakeData(structure reflect.Type) reflect.Value {
	new_value := reflect.New(structure).Interface()

	err := faker.FakeData(&new_value)

	if err != nil {
		log.Fatalf("Panic %v: ", err)
	}

	return reflect.ValueOf(new_value).Elem()
}
