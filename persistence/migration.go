package persistence

import (
	"github.com/hervit0/barulho/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	//db.DropTableIfExists(&model.Video{}, &model.Artist{}, &model.City{}, &model.Song{})
	db.Debug().AutoMigrate(&model.Video{}, &model.Artist{}, &model.City{}, &model.Song{})

	db.Debug().Model(&model.Video{}).AddForeignKey("song_id_int", "songs (song_id)", "CASCADE", "CASCADE")
}
