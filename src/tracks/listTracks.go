package tracks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) ListTracks(c *gin.Context) {
	tracks, err := handler.Queries.GetAllTracks(c)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, tracks)
}
