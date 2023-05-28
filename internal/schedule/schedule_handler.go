package schedule

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *handler) ScheduleFeed(c *gin.Context) {
	var scheduledFeedRequest ScheduleFeedRequest
	if err := c.ShouldBindJSON(&scheduledFeedRequest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(scheduledFeedRequest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	res, err := h.service.AddSchedule(c.Request.Context(), &scheduledFeedRequest)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  res,
	})
}

func (h *handler) DeviceScheduleList(c *gin.Context) {
	id := c.Param("id")
	deviceId, err := utils.StringToUUID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	res, err := h.service.ShowDeviceSchedules(c.Request.Context(), *deviceId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  res,
	})
}

func (h *handler) ScheduleList(c *gin.Context) {
	res, err := h.service.ShowAllSchedules(c.Request.Context())
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  res,
	})
}
