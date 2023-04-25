package main

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/connections"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/routes"
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
	defer mqtt.CloseConnection()

	pg, err := connections.NewPostgresDB()
	if err != nil {
		panic("Failed to connect to postgres database : " + err.Error())
	}
	defer pg.CloseConnection()

	routes.InitRoutes(mqtt, pg)
	routes.StartServer()
}
