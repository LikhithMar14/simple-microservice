package http

import (
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserId      string           `json:"userId"`
	PickUp      types.Coordinate `json:"pickUp"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(c *gin.Context) {
	log.Println("Inside the trip service")
	log.Println("HELLO")
	var req previewTripRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	trip, err := h.Service.GetRoute(c, &req.PickUp, &req.Destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create trip"})
		return
	}

	c.JSON(http.StatusOK, trip)
}
