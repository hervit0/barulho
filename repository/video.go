package repository

import (
	"fmt"
	"github.com/hervit0/barulho/model"
	"github.com/jinzhu/gorm"
)

type Video interface {
	FindBySongName(songName string) []VideoResult
	FindByCityName(cityName string) []VideoResult
	FindByArtistName(cityName string) []VideoResult
}

type VideoImpl struct {
	Db *gorm.DB
}

type VideoResult struct {
	VideoUid   string
	CityName   string
	ArtistName string
	SongName   string
	SongId     int64
}

//http://gorm.io/docs/query.html
func (video *VideoImpl) FindBySongName(songName string) []VideoResult {
	// TODO: Bulletproof that
	var songs []model.Song
	request := fmt.Sprintf("%%%v%%", songName)
	video.Db.Where("Title LIKE ?", request).Find(&songs)

	results := make([]VideoResult, len(songs))
	for i, song := range songs {
		results[i] = VideoResult{
			CityName:   video.getCityName(song.CityId),
			ArtistName: video.getArtistName(song.ArtistId),
			SongName:   song.Title,
			SongId:     song.SongId,
		}
	}
	return results
}

func (video *VideoImpl) FindByCityName(cityName string) []VideoResult {
	city := video.getCityByName(&cityName)

	var songs []model.Song
	video.Db.Where("city_id = ?", city.Id).Find(&songs)

	results := make([]VideoResult, len(songs))
	for i, song := range songs {
		results[i] = VideoResult{
			CityName:   video.getCityName(song.CityId),
			ArtistName: video.getArtistName(song.ArtistId),
			SongName:   song.Title,
			SongId:     song.SongId,
		}
	}
	return results
}

func (video *VideoImpl) FindByArtistName(artistName string) []VideoResult {
	artist := video.getArtistByName(&artistName)

	var songs []model.Song
	video.Db.Where("artist_id = ?", artist.Id).Find(&songs)

	results := make([]VideoResult, len(songs))
	for i, song := range songs {
		results[i] = VideoResult{
			CityName:   video.getCityName(song.CityId),
			ArtistName: video.getArtistName(song.ArtistId),
			SongName:   song.Title,
			SongId:     song.SongId,
		}
	}
	return results
}

func (video *VideoImpl) getCityByName(cityName *string) *model.City {
	var city model.City
	request := fmt.Sprintf("%%%v%%", *cityName)
	video.Db.Where("Title LIKE ?", request).Find(&city)

	return &city
}

func (video *VideoImpl) getArtistByName(artistName *string) *model.Artist {
	var artist model.Artist
	request := fmt.Sprintf("%%%v%%", *artistName)
	video.Db.Where("Title LIKE ?", request).Find(&artist)

	return &artist
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
