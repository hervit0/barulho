package models

type Video struct {
	VideoUid string `json:"video_uid",gorm:"primary_key"`
	Song     *Song  `json:"song"`
}
