package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	httpAddr = ":8083"
)

func main() {
	log.Println("Starting Trip Service")
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	httphandler := h.HttpHandler{Service: svc}

	r := gin.Default()
	log.Println("Trip Service listening on port 8083")
	r.POST("/preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Trip Service listening on %s", httpAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Server error occurred in trip service: %v", err)
	case sig := <-shutdown:
		log.Printf("Shutting down Trip Service due to %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
			log.Println("Forcing server close...")
			_ = server.Close()
		} else {
			log.Println("Server shut down gracefully.")
		}
	}
}
