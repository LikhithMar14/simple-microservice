package domain

import (
	"context"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID `json:"id"`
	UserId   string              `json:"userId" binding:"required"`
	Status   string              `json:"status" binding:"required"`
	RideFare *RideFareModel      `json:"rideFare" binding:"required,dive"`
}


type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel)(*TripModel, error)

}
type TripService interface {
	CreateTrip(ctx context.Context, trip *RideFareModel)(*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination *types.Coordinate)(*types.OsrmApiResponse,error)
}