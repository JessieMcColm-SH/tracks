package tracks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) ListOrigins(c *gin.Context) {
	tags, err := handler.Queries.GetAllOrigins(c)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, tags)
}
