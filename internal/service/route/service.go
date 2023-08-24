package route

import (
	"context"

	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/domain"
	"gopkg.in/guregu/null.v4"
)

type Service interface {
	GetRoutes(ctx context.Context) (routes domain.RouteResponse, err error)
	CreateRoute(ctx context.Context) (err error)
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger, db *sqlx.DB) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) GetRoutes(ctx context.Context) (routes domain.RouteResponse, err error) {
	routes = domain.RouteResponse{
		Name: null.StringFrom("Teste"),
	}
	return
}

func (s *service) CreateRoute(ctx context.Context) (err error) {
	return
}
