package resolver

import (
	"context"
	"github.com/hervit0/barulho/repository"
)

type Main interface {
	Hello() string
	GetVideosBySongName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error)
}

type MainImpl struct {
	VideosRepo repository.Video
}

func (_ *MainImpl) Hello() string { return "Hello from MainImpl" }

func (main *MainImpl) GetVideosBySongName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error) {
	results := main.VideosRepo.FindBySongName(args.Name)
	songResolvers := make([]*SongResolver, len(results))
	for i, song := range results {
		songResolvers[i] = &SongResolver{Song: song}
	}

	return &songResolvers, nil
}
