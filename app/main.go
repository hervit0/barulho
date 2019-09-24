package main

import (
	"encoding/json"
	"github.com/hervit0/barulho/models"
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

	for _, v := range videos {
		log.Printf("%+v", v.VideoUid)
		if v.Song != nil {
			log.Printf("%+v", *v.Song)
		}
	}
}
