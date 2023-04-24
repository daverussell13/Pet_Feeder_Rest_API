package main

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/cmd/api/routes"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/connections"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/feeder"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mqtt, err := connections.NewMqtt()
	if err != nil {
		panic("Failed to connect to mqtt broker : " + err.Error())
	}

	feederService := feeder.NewService(mqtt)
	feederHandler := feeder.NewHandler(feederService)

	handlers := routes.ApiHandlers{
		V1: routes.ApiV1Handlers{
			Feeder: feederHandler,
		},
	}

	routes.InitRoutes(handlers)
	routes.StartServer()

	defer mqtt.CloseConnection()
}
