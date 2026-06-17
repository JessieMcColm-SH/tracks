package main

import (
	"database/sql"
	"log"
	tracks "tracks/src/tracks"
	database "tracks/src/tracks/database"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func closeDB(database *sql.DB) {
	err := database.Close()
	log.Println(err)
}

func main() {
	log.Println("hello")
	DB := database.InitDB()
	defer closeDB(DB)
	rabbitQueries := database.NewTrackQueries(DB)
	trackHandler := tracks.NewTrackHandler(rabbitQueries)
	router := gin.Default()

	//router.GET("/bunnies", trackHandler.ListTracks)
	//router.GET("/bunnies/:id", trackHandler.GetRabbit)
	//router.DELETE("/bunnies/:id", trackHandler.DeleteRabbit)
	router.POST("/tracks", trackHandler.CreateTrack)
	//router.POST("/bunnies/:id", trackHandler.UpdateRabbit)

	log.Fatal(router.Run("localhost:8080"))
}
