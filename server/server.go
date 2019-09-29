package server

import (
	"github.com/hervit0/barulho/resolver"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Server struct {
	Resolver resolver.Main
}

func (server *Server) Do() {
	schema := server.parseSchema()
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (server *Server) parseSchema() *graphql.Schema {
	schemaFile, err := os.Open("./schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer schemaFile.Close()

	byteSchema, err := ioutil.ReadAll(schemaFile)
	if err != nil {
		log.Fatal(err)
	}
	return graphql.MustParseSchema(string(byteSchema), server.Resolver, graphql.UseFieldResolvers())
}
