package api

import (
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/realtime"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/schedule"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func InitRoutes(mqtt *mqtt.Mqtt, db *sql.DB) *gin.Engine {
	handler := initHandler(mqtt, db)
	return setupRoutes(handler)
}

func initHandler(mqtt *mqtt.Mqtt, db *sql.DB) *Handlers {
	// Realtime handler
	realtimeService := realtime.NewService(mqtt)
	realtimeHandler := realtime.NewHandler(realtimeService)

	// Schedule handler
	scheduleRepository := schedule.NewRepository(db)
	scheduleService := schedule.NewService(scheduleRepository, mqtt)
	scheduleHandler := schedule.NewHandler(scheduleService)

	v1 := &V1Handlers{
		realtime: realtimeHandler,
		schedule: scheduleHandler,
	}

	return NewHandler(v1)
}

func setupRoutes(handlers *Handlers) *gin.Engine {
	r := gin.Default()

	// Init routes custom validator
	initValidator()

	// Test default connection
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Api V1 routes
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/realtime", handlers.V1.realtime.RealtimeFeed)
	apiV1.POST("/schedule", handlers.V1.schedule.ScheduleFeed)
	apiV1.GET("/schedule", handlers.V1.schedule.ScheduleList)
	apiV1.GET("/schedule/:id", handlers.V1.schedule.DeviceScheduleList)
	return r
}

func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("feedTimeFormat", FeedTimeFormatValidator)
		if err != nil {
			return
		}
	}
}
