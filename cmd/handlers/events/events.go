package events

import (
	"AnthonyNixon/car-mileage-tracker/cmd/services/events"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEventsV1(router *gin.Engine) {
	router.POST("/v1/events", PostEvent)
}

func PostEvent(c *gin.Context) {
	var newEvent events.Event
	err := c.ShouldBindJSON(&newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad JSON Input, could not bind.",
			"Details": err.Error(),
		})
		return
	}

	error := events.NewEvent(newEvent)
	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
