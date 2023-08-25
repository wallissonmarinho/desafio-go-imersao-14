package service

import (
	"database/sql"

	"github.com/go-kit/log"

	Route "github.com/wallissonmarinho/desafio-go-imersao-14/internal/service/route"
)

type ServiceFactory interface {
	Route() Route.Service
}

type serviceFactory struct {
	RouteService Route.Service
}

func NewServiceFactory(db *sql.DB, logger log.Logger) ServiceFactory {
	return &serviceFactory{
		RouteService: Route.NewService(logger, db),
	}
}

func (s *serviceFactory) Route() Route.Service {
	return s.RouteService
}
