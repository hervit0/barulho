package main

import (
	"encoding/json"
	"github.com/hervit0/barulho/models"
	"github.com/hervit0/barulho/persistence"
	"github.com/hervit0/barulho/repository"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, _ := os.Open("./resource/video_data_full.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var videos []models.Video
	_ = json.Unmarshal(byteValue, &videos)

	db := persistence.Connect()
	persistence.Migrate(db)

	//for _, v := range videos {
	//	db.Save(&v)
	//	if v.Song != nil {
	//		db.Save(v.Song)
	//		if v.Song.Artist != nil {
	//			db.Save(v.Song.Artist)
	//		}
	//		if v.Song.City != nil {
	//			db.Save(v.Song.City)
	//		}
	//	}
	//}
	videosRepo := repository.Video{
		Db: db,
	}
	videosResult := videosRepo.FindBySongName("Ranchera")
	log.Printf("%+v", videosResult)
}
