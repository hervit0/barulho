package models

type Artist struct {
	Id    int64  `json:"id" gorm:"unique;primary_key"`
	Title string `json:"title"`
}
