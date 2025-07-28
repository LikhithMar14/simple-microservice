package main

import (
	"log"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("INSIDE THE TRIP SERVICE")
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	httphandler := h.HttpHandler{Service: svc}

	router := gin.Default()

	router.POST("preview", httphandler.HandleTripPreview)

	if err := router.Run(":8083"); err != nil {
		panic(err)

	}
}
