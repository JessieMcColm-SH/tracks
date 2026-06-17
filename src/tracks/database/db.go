package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	createTableQuery := `CREATE TABLE IF NOT EXISTS origins (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	origin TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS artists (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	artist TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS tracks (
	id TEXT PRIMARY KEY,
	creation_date TEXT NOT NULL,
	track_location TEXT NOT NULL,
	track_title TEXT NOT NULL,
	artist_id INTEGER NOT NULL,
	origin_id INTEGER NOT NULL,
	FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
	FOREIGN KEY (origin_id) REFERENCES origins (id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	tag TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS track_tags (
	track_id TEXT NOT NULL,
	tag_id TEXT NOT NULL,
	PRIMARY KEY (track_id, tag_id),
	FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
	);
	INSERT OR IGNORE INTO artists (artist) VALUES
	("No artist provided");

	INSERT OR IGNORE INTO origins (origin) VALUES
	("No origin provided");
	END
	`
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	if _, err := tx.Exec(createTableQuery); err != nil {
		log.Println("error in exec")
		log.Fatal(err)
	}
	tx.Commit()

	return db
}
