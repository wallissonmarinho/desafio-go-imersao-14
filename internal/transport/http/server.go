package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/endpoint"
)

type server struct {
	endpoint *endpoint.Endpoints
	logger   *log.Logger
}

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(context context.Context, db *sqlx.DB, endpoint *endpoint.Endpoints, logger *log.Logger) http.Handler {
	rest := &server{
		endpoint: endpoint,
		logger:   logger,
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(CorsMiddleware())

	r.Get("/api/routes", rest.GetRoutes)
	r.Post("/api/routes", rest.CreateRoute)

	return r
}

func CorsMiddleware() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
