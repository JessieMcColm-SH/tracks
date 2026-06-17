package tracks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *TrackHandler) DeleteTag(c *gin.Context) {
	err := handler.Queries.DeleteTag(c, c.Param("id"))
	if err != nil {
		//prob want bad request if no matching record?
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}
