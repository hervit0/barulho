package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func Do() {
	schema := parseSchema()
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseSchema() *graphql.Schema {
	schemaFile, _ := os.Open("./server/schema.graphql")
	defer schemaFile.Close()

	byteSchema, err := ioutil.ReadAll(schemaFile)
	if err != nil {
		log.Fatal(err)
	}
	return graphql.MustParseSchema(string(byteSchema), &query{})
}
