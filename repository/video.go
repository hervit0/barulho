package repository

import (
	"github.com/hervit0/barulho/models"
	"github.com/jinzhu/gorm"
	"log"
)

type Video struct {
	Db *gorm.DB
}

type VideoResponse struct {
	VideoUid   string
	CityName   string
	ArtistName string
	SongName   string
	SongId     int64
}

//http://gorm.io/docs/query.html
func (video *Video) FindBySongName(songName string) []VideoResponse {
	var songs []models.Song
	video.Db.Where("Title LIKE ?", songName).Find(&songs)
	log.Print(songs)

	videoResponses := make([]VideoResponse, len(songs))
	for i, song := range songs {
		videoResponses[i] = VideoResponse{
			CityName:   video.getCityName(song.CityId),
			ArtistName: video.getArtistName(song.ArtistId),
			SongName:   song.Title,
			SongId:     song.SongId,
		}
	}
	return videoResponses
}

func (video *Video) getCityName(cityId *int64) (cityName string) {
	if cityId == nil {
		return cityName
	}

	var city models.City
	video.Db.First(&city, *cityId)

	return city.Title
}

func (video *Video) getArtistName(artistId *int64) (artistName string) {
	if artistId == nil {
		return artistName
	}

	var artist models.Artist
	video.Db.First(&artist, *artistId)

	return artist.Title
}
