package main

import (
	"encoding/json"
	"github.com/hervit0/barulho/models"
	"github.com/hervit0/barulho/repository"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, _ := os.Open("./resource/video_data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var videos []models.Video
	_ = json.Unmarshal(byteValue, &videos)

	db := repository.Connect()
	db.DropTableIfExists(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})
	db.Debug().AutoMigrate(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})

	db.Debug().Model(&models.Video{}).AddForeignKey("song_id_int", "songs (song_id)", "CASCADE", "CASCADE")
	//db.Debug().Model(&models.Song{}).AddForeignKey("artist_id", "artists (id)", "CASCADE", "CASCADE")
	//db.Debug().Model(&models.Song{}).AddForeignKey("city_id", "cities (id)", "CASCADE", "CASCADE")

	for _, v := range videos {
		db.Debug().Save(&v)
	}

	var video models.Video
	var song models.Song
	var artist models.Artist
	var city models.City
	db.First(&video)
	db.First(&song)
	db.First(&artist)
	db.First(&city)

	log.Printf("%+v", video)
	log.Printf("%+v", *video.SongIdInt)
	log.Printf("%+v", song)
	log.Printf("%+v", artist)
	log.Printf("%+v", city)
}
