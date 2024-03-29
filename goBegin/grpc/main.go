package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

type query struct {
}

func (_ *query) Hello() string {
	return "Hello, world!"
}

func main() {
	s := `
		type Query {
			hello: String!
		}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
