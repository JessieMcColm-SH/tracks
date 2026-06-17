package tracks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) GetArtist(c *gin.Context) {
	tag, err := handler.Queries.GetArtist(c, c.Param("id"))
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, tag)
}
