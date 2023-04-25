package feeder

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type handler struct {
	feederService Service
}

func NewHandler(feeder Service) Handler {
	return &handler{
		feederService: feeder,
	}
}

func (h *handler) RealtimeFeed(c *gin.Context) {
	var realtimeFeedRequest RealtimeFeedRequest
	if err := c.ShouldBindJSON(&realtimeFeedRequest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	res, err := h.feederService.RealtimeFeed(c.Request.Context(), &realtimeFeedRequest)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  res,
	})
}
