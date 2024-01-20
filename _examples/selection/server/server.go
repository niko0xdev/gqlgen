package main

import (
	"log"
	"net/http"

	"github.com/niko0xdev/gqlgen/_examples/selection"
	"github.com/niko0xdev/gqlgen/graphql/handler"
	"github.com/niko0xdev/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
