package internal

import (
	"context"
	"fmt"
	models "tracks/src/tracks/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

type TrackStore interface {
	getAll(ctx context.Context) ([]models.Track, error)
}
