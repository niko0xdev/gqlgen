package graphql

import (
	"log"
	"os"
)

func GetVersion() string {
	content, err := os.ReadFile(string("../../VERSION"))
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
