package models

type Song struct {
	Id       int64   `json:"id"`
	Title    string  `json:"title"`
	ArtistId *int64  `json:"artist_id"`
	CityId   *int64  `json:"city_id"`
	Artist   *Artist `json:"artist"`
	City     *City   `json:"city"`
}
