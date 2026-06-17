package models

import (
	"time"
)

type Track struct {
	ID           int       `json:"id"`
	CreationDate time.Time `json:"creation_date"`
	FileLocation string    `json:"file_location"`
	TrackTitle   string    `json:"track_title"`
	ArtistId     string    `json:"artist_id"`
	OriginId     string    `json:"origin_id"`
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
