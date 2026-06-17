package internal

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	models "tracks/src/tracks/models"
)

type TrackQueries struct {
	DB *sql.DB
}

func NewTrackQueries(db *sql.DB) *TrackQueries {
	return &TrackQueries{DB: db}
}

func (r *TrackQueries) CreateTrack(ctx context.Context, track models.Track) (string, error) {
	id := uuid.New()
	sqlStatement := `
	INSERT INTO tracks (id, creation_date, track_location, track_title, artist_id, origin_id)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.ExecContext(ctx, sqlStatement, id.String(), track.CreationDate, track.TrackLocation, track.TrackTitle, track.ArtistId, track.OriginId)
	if err != nil {
		log.Println("failed to insert")
		log.Fatal(err.Error())
		return "", err
	}
	return id.String(), err
}

func (r *TrackQueries) CreateTag(ctx context.Context, tag models.Tag) (int, error) {
	id := 0
	sqlStatement := `
	INSERT INTO tags (tag)
	VALUES ($1)
	RETURNING id`
	err := r.DB.QueryRowContext(ctx, sqlStatement, tag.TagName).Scan(&id)
	if err != nil {
		log.Println("failed to insert")
		log.Print(err.Error())
		return 0, err
	}
	return id, err
}

func (r *TrackQueries) CreateArtist(ctx context.Context, artist models.Artist) (int, error) {
	id := 0
	sqlStatement := `
	INSERT INTO artists (artist)
	VALUES ($1)
	RETURNING id`
	err := r.DB.QueryRowContext(ctx, sqlStatement, artist.Artist).Scan(&id)
	if err != nil {
		log.Println("failed to insert")
		log.Print(err.Error())
		return 0, err
	}
	return id, err
}

func (r *TrackQueries) CreateOrigin(ctx context.Context, origin models.Origin) (int, error) {
	id := 0
	sqlStatement := `
	INSERT INTO origins (origin)
	VALUES ($1)
	RETURNING id`
	err := r.DB.QueryRowContext(ctx, sqlStatement, origin.Origin).Scan(&id)
	if err != nil {
		log.Println("failed to insert")
		log.Print(err.Error())
		return 0, err
	}
	return id, err
}

func (r *TrackQueries) UpdateTrack(ctx context.Context, id string, track models.Track) error {
	sqlStatement := `
	UPDATE tracks
	SET creation_date = $2, track_location = $3, track_title = $4, artist_id = $5, origin_id = $6
	WHERE  id = $1`
	_, err := r.DB.ExecContext(ctx, sqlStatement, id, track.CreationDate, track.TrackLocation, track.TrackTitle, track.ArtistId, track.OriginId)
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

func (r *TrackQueries) DeleteTag(ctx context.Context, id string) error {
	sqlStatement := `
	DELETE FROM tags
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
	SELECT id, creation_date, track_location, track_title, artist_id, origin_id FROM tracks
	WHERE id = $1 `
	row, err := r.DB.QueryContext(ctx, sqlStatement, id)
	if err != nil {
		return track, err
	}
	row.Next()
	err = row.Scan(&track.ID, &track.CreationDate, &track.TrackLocation, &track.TrackTitle, &track.ArtistId, &track.OriginId)
	return track, err
}

func (r *TrackQueries) GetTag(ctx context.Context, id string) (models.Tag, error) {
	var tag models.Tag
	sqlStatement := `
	SELECT id, tag FROM tags
	WHERE id = $1 `
	row, err := r.DB.QueryContext(ctx, sqlStatement, id)
	if err != nil {
		return tag, err
	}
	row.Next()
	err = row.Scan(&tag.ID, &tag.TagName)
	return tag, err
}

func (r *TrackQueries) GetAllTracks(ctx context.Context) ([]models.Track, error) {
	var tracks = []models.Track{}
	rows, err := r.DB.Query("SELECT id, creation_date, track_location, track_title, artist_id, origin_id FROM tracks")

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
		err = rows.Scan(&track.ID, &track.CreationDate, &track.TrackLocation, &track.TrackTitle, &track.ArtistId, &track.OriginId)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tracks, nil
}

func (r *TrackQueries) GetAllTags(ctx context.Context) ([]models.Tag, error) {
	var tags = []models.Tag{}
	rows, err := r.DB.Query("SELECT id, tag FROM tags")

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
		var tag models.Tag
		err = rows.Scan(&tag.ID, &tag.TagName)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tags, nil
}

func (r *TrackQueries) GetAllArtists(ctx context.Context) ([]models.Artist, error) {
	var artists = []models.Artist{}
	rows, err := r.DB.Query("SELECT id, artist FROM artists")

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
		var artist models.Artist
		err = rows.Scan(&artist.ID, &artist.Artist)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return artists, nil
}

func (r *TrackQueries) GetAllOrigins(ctx context.Context) ([]models.Origin, error) {
	var origins = []models.Origin{}
	rows, err := r.DB.Query("SELECT id, origin FROM origins")

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
		var origin models.Origin
		err = rows.Scan(&origin.ID, &origin.Origin)
		if err != nil {
			return nil, err
		}
		origins = append(origins, origin)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return origins, nil
}
