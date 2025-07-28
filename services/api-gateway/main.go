package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"

	"github.com/gin-gonic/gin"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Ride Sharing API Gateway",
		})
	})
	r.POST("/trip/preview",handleTripPreview)
	server := &http.Server {
		Addr: httpAddr,
		Handler: r,
	}
 
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error occured in gateway ",err)
	}
}
