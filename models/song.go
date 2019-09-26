package models

type Song struct {
	//gorm.Model
	SongId   int64  `json:"id" gorm:"unique;primary_key"`
	Title    string `json:"title"`
	VideoUid *string
	ArtistId *int64  `json:"artist_id"`
	CityId   *int64  `json:"city_id"`
	Artist   *Artist `json:"artist",gorm:"foreignkey:ArtistId;association_foreignkey:ArtistId"`
	City     *City   `json:"city",gorm:"foreignkey:CityId;association_foreignkey:CityId"`
}
