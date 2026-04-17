package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	res := HealthResponse{
		Status: "OK",
	}

	c.JSON(http.StatusOK, res)
}
