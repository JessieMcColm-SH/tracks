package tracks

import (
	"net/http"
	models "tracks/src/tracks/models"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) CreateArtist(c *gin.Context) {
	var artist models.Artist
	if err := c.BindJSON(&artist); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	id, err := handler.Queries.CreateArtist(c, artist)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	artist.ID = id
	c.IndentedJSON(http.StatusCreated, artist)
}
