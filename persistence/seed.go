package persistence

import (
	"encoding/json"
	"github.com/hervit0/barulho/model"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
)

func Seed(db *gorm.DB) {
	jsonFile, _ := os.Open("./resource/video_data_full.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var videos []model.Video
	_ = json.Unmarshal(byteValue, &videos)

	for _, v := range videos {
		db.Debug().Save(&v)
		if v.Song != nil {
			db.Debug().Save(v.Song)
			if v.Song.Artist != nil {
				db.Debug().Save(v.Song.Artist)
			}
			if v.Song.City != nil {
				db.Debug().Save(v.Song.City)
			}
		}
	}
}
