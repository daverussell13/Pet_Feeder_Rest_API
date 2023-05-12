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
	"os"
)

var r *gin.Engine

func StartServer(mqtt *mqtt.Mqtt, db *sql.DB) {
	handler := initHandler(mqtt, db)
	initRoutes(handler)
	runServer()
}

func initHandler(mqtt *mqtt.Mqtt, db *sql.DB) *Handlers {
	// Realtime handler
	realtimeService := realtime.NewService(mqtt)
	realtimeHandler := realtime.NewHandler(realtimeService)

	// Schedule handler
	scheduleRepository := schedule.NewRepository(db)
	scheduleService := schedule.NewService(scheduleRepository)
	scheduleHandler := schedule.NewHandler(scheduleService)

	v1 := &V1Handlers{
		realtime: realtimeHandler,
		schedule: scheduleHandler,
	}

	return NewHandler(v1)
}

func initRoutes(handlers *Handlers) {
	r = gin.Default()

	initValidator()

	// Test connection
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/realtime", handlers.V1.realtime.RealtimeFeed)
	apiV1.POST("/schedule", handlers.V1.schedule.ScheduledFeed)
	apiV1.GET("/schedule", handlers.V1.schedule.ScheduleList)
}

func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timeFormat", TimeFormatValidator)
		if err != nil {
			return
		}
	}
}

func runServer() {
	serverAddress := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	if err := r.Run(serverAddress); err != nil {
		panic("Couldn't start api : " + err.Error())
	}
}
