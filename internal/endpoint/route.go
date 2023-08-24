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

		_ = level.Error(logger).Log("message", "ok")

		return nil, nil
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
