package tracks

import (
	"log"
	"net/http"
	models "tracks/src/tracks/models"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) UpdateTrack(c *gin.Context) {
	var track models.Track
	if err := c.BindJSON(&track); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	err := handler.Queries.UpdateTrack(c, c.Param("id"), track)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	track, err = handler.Queries.GetTrack(c, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusCreated, track)
}
