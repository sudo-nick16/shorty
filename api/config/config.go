package config

import (
	"log"
	"os"
	"reflect"
)

type Env struct {
    MONGO_URI string
    PORT string
    SERVER string
}

var E Env

func LoadConfig() {
	E.MONGO_URI = os.Getenv("MONGO_URI")
	E.SERVER = os.Getenv("SERVER")
	E.PORT = os.Getenv("PORT")

	fields := reflect.TypeOf(E)
	values := reflect.ValueOf(E)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if value.String() == "" {
			log.Fatalf("Missing environment variable: %v", field.Name)
		}
	}
}

     
     
     
