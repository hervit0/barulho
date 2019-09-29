package resolver

import (
	"context"
	"github.com/hervit0/barulho/repository"
)

type Resolver interface {
	Hello() string
	GetVideosBySongName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error)
	GetVideosByCityName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error)
	GetVideosByArtistName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error)
}

type ResolverImpl struct {
	VideosRepo repository.Video
}

func (_ *ResolverImpl) Hello() string { return "Hello from ResolverImpl" }

func (main *ResolverImpl) GetVideosBySongName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error) {
	results := main.VideosRepo.FindBySongName(args.Name)
	songResolvers := make([]*SongResolver, len(results))
	for i, song := range results {
		songResolvers[i] = &SongResolver{Song: song}
	}

	return &songResolvers, nil
}

func (main *ResolverImpl) GetVideosByCityName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error) {
	results := main.VideosRepo.FindByCityName(args.Name)
	songResolvers := make([]*SongResolver, len(results))
	for i, song := range results {
		songResolvers[i] = &SongResolver{Song: song}
	}

	return &songResolvers, nil
}

func (main *ResolverImpl) GetVideosByArtistName(ctx context.Context, args struct{ Name string }) (*[]*SongResolver, error) {
	results := main.VideosRepo.FindByArtistName(args.Name)
	songResolvers := make([]*SongResolver, len(results))
	for i, song := range results {
		songResolvers[i] = &SongResolver{Song: song}
	}

	return &songResolvers, nil
}
