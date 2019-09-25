package models

type Song struct {
	Id       int64   `json:"id",gorm:"primary_key"`
	Title    string  `json:"title"`
	ArtistId *int64  `json:"artist_id"`
	CityId   *int64  `json:"city_id"`
	Artist   *Artist `json:"artist",gorm:"foreignkey:ArtistId"`
	City     *City   `json:"city",gorm:"foreignkey:CityId"`
}
