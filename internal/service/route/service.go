package route

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"

	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/domain"
	routeRepo "github.com/wallissonmarinho/desafio-go-imersao-14/internal/repository/routes"
)

type Service interface {
	GetRoutes(ctx context.Context) (routes []domain.RouteResponse, err error)
	CreateRoute(ctx context.Context, request domain.RouteRequest) (route domain.RouteResponse, err error)
}

type service struct {
	logger    log.Logger
	routeRepo routeRepo.RoutesRepositoryInterface
}

func NewService(logger log.Logger, db *sql.DB) Service {
	return &service{
		logger:    logger,
		routeRepo: routeRepo.NewRoutesRepository(db),
	}
}

func (s *service) GetRoutes(ctx context.Context) (routeResponse []domain.RouteResponse, err error) {

	routes, err := s.routeRepo.FindAll()
	if err != nil {
		return
	}

	for _, route := range routes {
		routeResponse = append(routeResponse, domain.RouteResponse{
			ID:   route.ID,
			Name: route.Name,
			Source: domain.Coordinates{
				Lat: route.SourceLat,
				Lng: route.SourceLng,
			},
			Destination: domain.Coordinates{
				Lat: route.DestLat,
				Lng: route.DestLng,
			},
		})
	}

	return
}

func (s *service) CreateRoute(ctx context.Context, request domain.RouteRequest) (routeResponse domain.RouteResponse, err error) {

	routeID, err := s.routeRepo.CreateRoute(routeRepo.RouteRepo{
		Name:      request.Name,
		SourceLat: request.Source.Lat,
		SourceLng: request.Source.Lng,
		DestLat:   request.Destination.Lat,
		DestLng:   request.Destination.Lng,
	})
	if err != nil {
		return
	}

	route, err := s.routeRepo.GetRoute(routeID)
	if err != nil {
		return
	}

	routeResponse = domain.RouteResponse{
		ID:   route.ID,
		Name: route.Name,
		Source: domain.Coordinates{
			Lat: route.SourceLat,
			Lng: route.SourceLng,
		},
		Destination: domain.Coordinates{
			Lat: route.DestLat,
			Lng: route.DestLng,
		},
	}
	return
}
