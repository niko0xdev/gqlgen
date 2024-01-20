package main

import (
	"log"
	"net/http"

	"github.com/niko0xdev/gqlgen/_examples/scalars"
	"github.com/niko0xdev/gqlgen/graphql/handler"
	"github.com/niko0xdev/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
