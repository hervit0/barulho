package main

import (
	"encoding/json"
	"github.com/hervit0/barulho/models"
	"github.com/hervit0/barulho/repository"
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

	//for _, v := range videos {
	//	log.Printf("%+v", v.VideoUid)
	//	if v.Song != nil {
	//		log.Printf("%+v", *v.Song)
	//	}
	//}

	db := repository.Connect()
	db.Debug().DropTableIfExists(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})
	//db.Debug().Model(&models.Song{})
	//db.Debug().Model(&models.Video{}).Related(&models.Song{}, "Song")
	db.Debug().AutoMigrate(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})

	for _, v := range videos {
		if v.Song != nil {
			if v.Song.Artist != nil {
				db.FirstOrCreate(&v.Song.Artist, v.Song.Artist)
			}
			if v.Song.City != nil {
				db.FirstOrCreate(&v.Song.City, v.Song.City)
			}
			v.Song.VideoUid = &v.VideoUid
			db.Debug().Create(v.Song)
		}
		db.Debug().Create(v)
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
	log.Printf("%+v", song)
	log.Printf("%+v", artist)
	log.Printf("%+v", city)
}
