package models

type Track struct {
	ID            string `json:"id"`
	CreationDate  string `json:"creation_date"`
	TrackLocation string `json:"track_location"`
	TrackTitle    string `json:"track_title"`
	ArtistId      int    `json:"artist_id"`
	OriginId      int    `json:"origin_id"`
}

type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}

type Artist struct {
	ID     int    `json:"id"`
	Artist string `json:"artist"`
}

type Origin struct {
	ID     int    `json:"id"`
	Origin string `json:"origin"`
}
