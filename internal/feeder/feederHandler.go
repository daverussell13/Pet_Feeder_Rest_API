package feeder

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RealtimeFeed(c *gin.Context) {
	var realtimeFeedRequest RealtimeFeedRequest
	if err := c.ShouldBindJSON(&realtimeFeedRequest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	res, err := h.service.RealtimeFeed(c.Request.Context(), &realtimeFeedRequest)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
