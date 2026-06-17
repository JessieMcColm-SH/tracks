package tracks

import (
	"net/http"
	models "tracks/src/tracks/models"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) CreateOrigin(c *gin.Context) {
	var origin models.Origin
	if err := c.BindJSON(&origin); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	id, err := handler.Queries.CreateOrigin(c, origin)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	origin.ID = id
	c.IndentedJSON(http.StatusCreated, origin)
}
