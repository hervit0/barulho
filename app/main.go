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

	db.AutoMigrate(&models.Video{})
	db.AutoMigrate(&models.Artist{})
	db.AutoMigrate(&models.City{})

	db.Create(&models.City{
		Id:    22,
		Title: "Wow",
	})

	var city models.City
	db.First(&city)

	log.Printf("%+v", city)
}
