package models

type Video struct {
	//gorm.Model
	VideoUid string `json:"video_uid" gorm:"primary_key;unique"`
	Song     *Song  `json:"song" gorm:"ForeignKey:video_uid"`
	SongId   *int64
}
