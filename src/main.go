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

	router.GET("/tracks", trackHandler.ListTracks)
	router.GET("/tags", trackHandler.ListTags)
	router.GET("/artists", trackHandler.ListArtists)
	router.GET("/origins", trackHandler.ListOrigins)
	router.GET("/tracks/:id", trackHandler.GetTrack)
	router.GET("/tags/:id", trackHandler.GetTag)
	router.DELETE("/tracks/:id", trackHandler.DeleteTrack)
	router.DELETE("/tags/:id", trackHandler.DeleteTag)
	router.POST("/tracks", trackHandler.CreateTrack)
	router.POST("/tags", trackHandler.CreateTag)
	router.POST("/origins", trackHandler.CreateOrigin)
	router.POST("/artists", trackHandler.CreateArtist)
	//router.POST("/bunnies/:id", trackHandler.UpdateRabbit)

	log.Fatal(router.Run("localhost:8080"))
}
