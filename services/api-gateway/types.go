package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserId      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
