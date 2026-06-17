package tracks

import database "tracks/src/tracks/database"

type TrackHandler struct {
	Queries *database.TrackQueries
}

func NewTrackHandler(query *database.TrackQueries) *TrackHandler {
	return &TrackHandler{Queries: query}
}
