package main

import (
	"context"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/database"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/api"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mqttInstance, err := mqtt.NewMqtt()
	if err != nil {
		panic("Failed to connect to mqtt broker : " + err.Error())
	}
	defer mqttInstance.CloseConnection()

	pgInstance, err := database.NewPostgresDB()
	if err != nil {
		panic("Failed to connect to postgres database : " + err.Error())
	}
	defer pgInstance.CloseConnection()

	router := api.InitRoutes(mqttInstance, pgInstance.GetDB())

	server := &http.Server{
		Addr:    os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"),
		Handler: router,
	}

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
