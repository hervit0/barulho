package resolver

import (
	"context"
	"github.com/hervit0/barulho/repository"
)

type SongResolver struct {
	Song repository.VideoResult
}

func (t *SongResolver) Id(ctx context.Context) *int64 {
	return &t.Song.SongId
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
