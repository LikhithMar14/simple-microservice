package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRiderWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websockets upgrade error: %v", err)
		return
	}
	defer conn.Close()

	userID := c.Query("userID")

	if userID == "" {
		log.Printf("userID is empty")
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message %s ", err)
			break
		}
		log.Printf("Received message: %s", message)

	}

}

func handleDriverWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websockets upgrade error: %v", err)
		return
	}
	defer conn.Close()

	userID := c.Query("userID")
	if userID == "" {
		log.Println("No user ID is provided")
		return
	}

	packageSlug := c.Query("packageSlug")

	if packageSlug == "" {
		log.Printf("packageSlug is empty")
		return
	}

	type Driver struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		ProfilePicture string `json:"profilePicture"`
		CarPlate       string `json:"carPlate"`
		PackageSlug    string `json:"packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver{
			Id:             userID,
			Name:           "Likhith",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate:       "KA 01 AB 1234",
			PackageSlug:    packageSlug,
		},
	}
	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("Error writing message %s ", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Error reading message %s ", err)
			break
		}
		log.Printf("Received message: %s", message)
	}

}
