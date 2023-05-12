package main

import (
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/database"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mqttClient, err := mqtt.NewMqtt()
	if err != nil {
		panic("Failed to connect to mqtt broker : " + err.Error())
	}
	defer mqttClient.CloseConnection()

	pg, err := database.NewPostgresDB()
	if err != nil {
		panic("Failed to connect to postgres database : " + err.Error())
	}
	defer pg.CloseConnection()

	api.StartServer(mqttClient, pg.GetDB())
}
