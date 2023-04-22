package main

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	serverAddress := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")

	if err = router.Run(serverAddress); err != nil {
		log.Fatal("Couldn't start server")
	}
}
