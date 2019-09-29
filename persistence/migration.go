package persistence

import (
	"github.com/hervit0/barulho/models"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	//db.DropTableIfExists(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})
	db.Debug().AutoMigrate(&models.Video{}, &models.Artist{}, &models.City{}, &models.Song{})

	db.Debug().Model(&models.Video{}).AddForeignKey("song_id_int", "songs (song_id)", "CASCADE", "CASCADE")
}
