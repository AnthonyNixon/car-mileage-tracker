package fillups

import (
	"AnthonyNixon/car-mileage-tracker/cmd/services/fillups"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFillupsV1(router *gin.Engine) {
	router.POST("/v1/fillups", PostFillup)
}

func PostFillup(c *gin.Context) {
	var newFillup fillups.FillUp
	err := c.ShouldBindJSON(&newFillup)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad JSON Input, could not bind.",
			"Details": err.Error(),
		})
		return
	}

	error := fillups.NewFillup(newFillup)
	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
