package routes

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/connections"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/realtime"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/schedule"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

var r *gin.Engine

func StartServer(conn *connections.Connections) {
	// Realtime handler
	realtimeService := realtime.NewService(conn.Mqtt)
	realtimeHandler := realtime.NewHandler(realtimeService)

	// Schedule handler
	scheduleRepository := schedule.NewRepository(conn.PostgresDB.GetDB())
	scheduleService := schedule.NewService(scheduleRepository)
	scheduleHandler := schedule.NewHandler(scheduleService)

	v1 := &APIV1Handlers{
		realtime: realtimeHandler,
		schedule: scheduleHandler,
	}

	handlers := NewHandler(v1)

	setupRoutes(handlers)
	runServer()
}

func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timeFormat", TimeFormatValidator)
		if err != nil {
			return
		}
	}
}

func setupRoutes(handlers *Handlers) {
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

func runServer() {
	serverAddress := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	if err := r.Run(serverAddress); err != nil {
		panic("Couldn't start server : " + err.Error())
	}
}
