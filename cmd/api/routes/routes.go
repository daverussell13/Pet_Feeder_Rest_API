package routes

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/feeder"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type ApiV1Handlers struct {
	Feeder feeder.Handler
}

type ApiHandlers struct {
	V1 ApiV1Handlers
}

var r *gin.Engine

func InitRoutes(handlers ApiHandlers) {
	r = gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/feeder/realtime", handlers.V1.Feeder.RealtimeFeed)
}

func StartServer() {
	serverAddress := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	if err := r.Run(serverAddress); err != nil {
		panic("Couldn't start server : " + err.Error())
	}
}
