package tracks

import (
	"net/http"
	models "tracks/src/tracks/models"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	id, err := handler.Queries.CreateTag(c, tag)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	tag.ID = id
	c.IndentedJSON(http.StatusCreated, tag)
}
