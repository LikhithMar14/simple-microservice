package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}
func (s *service) CreateTrip(ctx context.Context, fareModel *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserId:   fareModel.UserId,
		Status:   "pending",
		RideFare: fareModel,
	}
	return s.repo.CreateTrip(ctx, t)
}
func (s *service) GetRoute(ctx context.Context, source, destination *types.Coordinate) (*types.OsrmApiResponse, error) {
	var routeResponse types.OsrmApiResponse

	url := fmt.Sprintf(
		"http://router.project-osrm.org/route/v1/driving/%.6f,%.6f;%.6f,%.6f?overview=full&geometries=geojson",
		source.Longitude, source.Latitude, destination.Longitude, destination.Latitude,
	)
	log.Print(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call OSRM API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OSRM API returned status %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&routeResponse); err != nil {
		return nil, fmt.Errorf("failed to decode OSRM response: %w", err)
	}

	if len(routeResponse.Routes) == 0 {
		return nil, fmt.Errorf("no route found")
	}

	return &routeResponse, nil
}

