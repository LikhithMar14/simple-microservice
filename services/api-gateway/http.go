package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
)

func handleTripPreview(c *gin.Context) {
	log.Println("handleTripPreview called")

	jsonBody, err := c.GetRawData()
	if err != nil {
		log.Println("ERROR reading body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to read request body"})
		return
	}

	log.Println("RAW DATA:", string(jsonBody))

	var req previewTripRequest
	if err := json.Unmarshal(jsonBody, &req); err != nil {
		log.Println("ERROR unmarshalling:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	log.Printf("Parsed request: %+v", req)
	tripServiceURL := "http://trip-service:8083/preview"

	resp, err := http.Post(tripServiceURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("ERROR calling trip-service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to contact trip service"})
		return
	}
	defer resp.Body.Close()

	log.Printf("Trip service responded with status: %d", resp.StatusCode)

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
	log.Println("Response sent to client")
}
