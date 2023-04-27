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

	connection := connections.NewConnection()

	routes.StartServer(connection)
	defer connection.CloseAllConnection()
}
