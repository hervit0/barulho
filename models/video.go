package models

type Video struct {
	//gorm.Model
	VideoUid string `json:"video_uid",gorm:"primary_key"`
	Song     *Song  `json:"song",gorm:"foreignkey:VideoUid"`
}
