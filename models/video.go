package models

type Video struct {
	VideoUid string `json:"video_uid"`
	Song     *Song  `json:"song"`
}
