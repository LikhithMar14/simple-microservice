package domain
import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID          primitive.ObjectID `json:"id"`
	UserId      string              `json:"userId" binding:"required"`
	PackageSlug string              `json:"packageSlug" binding:"required"`
	TotalPrice  float64             `json:"totalPrice" binding:"required,gt=0"`
}
