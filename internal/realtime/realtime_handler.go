package realtime

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/server"
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/utils"
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
	var realtimeFeedRequest FeedRequest
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
		switch err.Error() {
		case server.DeviceUnresponsive:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": utils.UcFirst(err.Error()),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  res,
	})
}
