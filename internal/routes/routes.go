package routes

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/connections"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/feeder"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/schedule"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"regexp"
)

var r *gin.Engine

type APIV1Handlers struct {
	feeder   feeder.Handler
	schedule schedule.Handler
}

func InitRoutes(mqtt *connections.Mqtt, pg *connections.PostgresDB) {
	feederService := feeder.NewService(mqtt)
	feederHandler := feeder.NewHandler(feederService)

	scheduleRepository := schedule.NewRepository(pg.GetDB())
	scheduleService := schedule.NewService(scheduleRepository)
	scheduleHandler := schedule.NewHandler(scheduleService)

	handlers := &APIV1Handlers{
		feeder:   feederHandler,
		schedule: scheduleHandler,
	}

	SetupRoutes(handlers)
}

func TimeFormatValidator(fl validator.FieldLevel) bool {
	timePattern := "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"
	regex := regexp.MustCompile(timePattern)
	timeStr := fl.Field().String()
	return regex.MatchString(timeStr)
}

func SetupRoutes(v1Hdl *APIV1Handlers) {
	r = gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("timeFormat", TimeFormatValidator)
		if err != nil {
			return
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/realtime", v1Hdl.feeder.RealtimeFeed)
	apiV1.POST("/schedule", v1Hdl.schedule.ScheduledFeed)
}

func StartServer() {
	serverAddress := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	if err := r.Run(serverAddress); err != nil {
		panic("Couldn't start server : " + err.Error())
	}
}
