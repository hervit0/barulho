package repository

import (
	"fmt"
	"github.com/hervit0/barulho/model"
	"github.com/jinzhu/gorm"
)

type Video interface {
	FindBySongName(songName string) []VideoResult
}

type VideoImpl struct {
	Db *gorm.DB
}

type VideoResult struct {
	VideoUid   string
	CityName   string
	ArtistName string
	SongName   string
	Id         int64
}

//http://gorm.io/docs/query.html
func (video *VideoImpl) FindBySongName(songName string) []VideoResult {
	var songs []model.Song
	request := fmt.Sprintf("%%%v%%", songName)
	video.Db.Where("Title LIKE ?", request).Find(&songs)

	results := make([]VideoResult, len(songs))
	for i, song := range songs {
		results[i] = VideoResult{
			CityName:   video.getCityName(song.CityId),
			ArtistName: video.getArtistName(song.ArtistId),
			SongName:   song.Title,
			Id:         song.SongId,
		}
	}
	return results
}

func (video *VideoImpl) getCityName(cityId *int64) (cityName string) {
	if cityId == nil {
		return cityName
	}

	var city model.City
	video.Db.First(&city, *cityId)

	return city.Title
}

func (video *VideoImpl) getArtistName(artistId *int64) (artistName string) {
	if artistId == nil {
		return artistName
	}

	var artist model.Artist
	video.Db.First(&artist, *artistId)

	return artist.Title
}
