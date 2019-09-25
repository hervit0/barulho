package models

type Artist struct {
	Id    int64  `json:"id",gorm:"primary_key"`
	Title string `json:"title"`
}
