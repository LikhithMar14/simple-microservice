package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/shared/env"

	"github.com/gin-gonic/gin"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")
	r := gin.Default()
	r.Use(enableCORS())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Ride Sharing API Gateway",
		})
	})
	r.POST("/trip/preview", handleTripPreview)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "api-gateway"})
	})
	log.Println("Registered route: POST /trip/preview")
	log.Println("Registered route: GET /health")

	// WebSocket routes
	ws := r.Group("/ws")
	{
		ws.GET("/drivers", handleDriverWebsocket)
		ws.GET("/riders", handleRiderWebsocket)
	}

	server := &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("API Gateway listening on %s", httpAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Error occurred in gateway: %v", err)

	case sig := <-shutdown:
		log.Printf("Shutting down API Gateway due to %v", sig)

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
