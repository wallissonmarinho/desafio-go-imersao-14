package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/service"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	CreateRouteEndpoint endpoint.Endpoint
	GetRoutesEndpoint   endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s service.ServiceFactory, logger log.Logger) Endpoints {
	return Endpoints{
		CreateRouteEndpoint: makeCreateRouteEndpoint(s, logger),
		GetRoutesEndpoint:   makeGetRoutesEndpoint(s, logger),
	}
}
