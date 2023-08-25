package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/domain"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/helpers"
)

func (s *server) CreateRoute(w http.ResponseWriter, r *http.Request) {
	var route domain.RouteRequest
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		helpers.JSON(w, http.StatusInternalServerError, nil)
		return
	}
	resp, err := s.endpoint.CreateRouteEndpoint(r.Context(), route)
	if err != nil {
		logrus.Error(err)
	}
	helpers.JSON(w, int(resp.(domain.Response).Code.Int64), resp.(domain.Response).Response)
}

func (s *server) GetRoutes(w http.ResponseWriter, r *http.Request) {
	resp, err := s.endpoint.GetRoutesEndpoint(r.Context(), nil)
	if err != nil {
		logrus.Error(err)
	}
	helpers.JSON(w, int(resp.(domain.Response).Code.Int64), resp.(domain.Response).Response)
}
