package tracks

import (
	"net/http"
	models "tracks/src/tracks/models"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) CreateTrack(c *gin.Context) {
	var track models.Track
	if err := c.BindJSON(&track); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	id, err := handler.Queries.CreateTrack(c, track)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	track.ID = id
	c.IndentedJSON(http.StatusCreated, track)
}
