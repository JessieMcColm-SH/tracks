package internal

import (
	"context"
	"database/sql"
	"log"

	models "tracks/src/tracks/models"
)

type TrackQueries struct {
	DB *sql.DB
}

func NewTrackQueries(db *sql.DB) *TrackQueries {
	return &TrackQueries{DB: db}
}

func (r *TrackQueries) CreateTrack(ctx context.Context, track models.Track) (int, error) {
	sqlStatement := `
	INSERT INTO tracks (creation_date, file_location, track_title, artist_id, origin_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := 0
	err := r.DB.QueryRowContext(ctx, sqlStatement, track.CreationDate, track.FileLocation, track.TrackTitle, track.ArtistId, track.OriginId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *TrackQueries) UpdateTrack(ctx context.Context, id string, track models.Track) error {
	sqlStatement := `
	UPDATE tracks
	SET creation_date = $2, file_location = $3, track_title = $4, artist_id = $5, origin_id = $6
	WHERE  id = $1`
	_, err := r.DB.ExecContext(ctx, sqlStatement, id, track.CreationDate, track.FileLocation, track.TrackTitle, track.ArtistId, track.OriginId)
	if err != nil {
		return err
	}
	return err
}

func (r *TrackQueries) DeleteTrack(ctx context.Context, id string) error {
	sqlStatement := `
	DELETE FROM tracks
	WHERE id = $1 `
	_, err := r.DB.ExecContext(ctx, sqlStatement, id)
	if err != nil {
		return err
	}
	return err
}

func (r *TrackQueries) GetTrack(ctx context.Context, id string) (models.Track, error) {
	var track models.Track
	sqlStatement := `
	SELECT id, creation_date, file_location, track_title, artist_id, origin_id FROM tracks
	WHERE id = $1 `
	row, err := r.DB.QueryContext(ctx, sqlStatement, id)
	if err != nil {
		return track, err
	}
	row.Next()
	err = row.Scan(&track.ID, &track.CreationDate, &track.FileLocation, &track.TrackTitle, &track.ArtistId, &track.OriginId)
	return track, err
}

func (r *TrackQueries) GetAll(ctx context.Context) ([]models.Track, error) {
	var tracks = []models.Track{}
	rows, err := r.DB.Query("SELECT id, creation_date, file_location, track_title, artist_id, origin_id FROM tracks")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		dErr := rows.Close()
		if dErr != nil && err == nil {
			err = dErr
		}
	}()
	for rows.Next() {
		var track models.Track
		err = rows.Scan(&track.ID, &track.CreationDate, &track.FileLocation, &track.TrackTitle, &track.ArtistId, &track.OriginId)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
