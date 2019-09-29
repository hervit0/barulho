package server

import (
	"context"
	"github.com/hervit0/barulho/repository"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type resolver struct{}

type SongResolver struct {
	Song repository.VideoResult
}

func (t *SongResolver) Id(ctx context.Context) *int64 {
	return &t.Song.Id
}

func (t *SongResolver) Name(ctx context.Context) *string {
	return &t.Song.SongName
}
func (t *SongResolver) CityName(ctx context.Context) *string {
	return &t.Song.CityName
}
func (t *SongResolver) ArtistName(ctx context.Context) *string {
	return &t.Song.ArtistName
}

func (_ *resolver) Hello() string { return "Hello from resolver" }

func (_ *resolver) GetSongByName(ctx context.Context, args struct{ Name string }) (*SongResolver, error) {
	result := repository.VideoResult{
		VideoUid:   "",
		CityName:   "Heaven",
		ArtistName: "Mock",
		SongName:   "Lol",
		Id:         322,
	}
	resolver := SongResolver{
		Song: result,
	}
	return &resolver, nil
}

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
	return graphql.MustParseSchema(string(byteSchema), &resolver{}, graphql.UseFieldResolvers())
}
