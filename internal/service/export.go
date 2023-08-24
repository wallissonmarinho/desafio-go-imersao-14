package service

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	Route "github.com/wallissonmarinho/desafio-go-imersao-14/internal/service/route"
)

type ServiceFactory interface {
	Route() Route.Service
}

type serviceFactory struct {
	RouteService Route.Service
}

func NewServiceFactory(db *sqlx.DB, logger log.Logger) ServiceFactory {
	return &serviceFactory{
		RouteService: Route.NewService(logger, db),
	}
}

func (s *serviceFactory) Route() Route.Service {
	return s.RouteService
}
