package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/domain"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/service"
	"gopkg.in/guregu/null.v4"
)

func makeCreateRouteEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if request == nil {
			return domain.Response{
				Code:     null.IntFrom(http.StatusBadRequest),
				Response: nil,
			}, nil
		}

		route := request.(domain.RouteRequest)
		resp, err := s.Route().CreateRoute(ctx, route)
		if err != nil {
			_ = level.Error(logger).Log("message", "invalid request")
			return domain.Response{
				Code:     null.IntFrom(http.StatusBadRequest),
				Response: err.Error(),
			}, nil
		}

		_ = level.Error(logger).Log("message", "ok")

		return domain.Response{
			Code:     null.IntFrom(http.StatusCreated),
			Response: resp,
		}, nil
	}
}

func makeGetRoutesEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		resp, err := s.Route().GetRoutes(ctx)
		if err != nil {
			_ = level.Error(logger).Log("message", "invalid request")
			return domain.Response{
				Code:     null.IntFrom(http.StatusBadRequest),
				Response: err.Error(),
			}, nil
		}

		_ = level.Error(logger).Log("message", "ok")

		return domain.Response{
			Code:     null.IntFrom(http.StatusOK),
			Response: resp,
		}, nil
	}
}
