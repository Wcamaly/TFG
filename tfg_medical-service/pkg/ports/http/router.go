package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"tfg/user-service/cmd/config"
	"tfg/user-service/pkg/handlers"
	"tfg/user-service/pkg/handlers/user"
)

// NewRouter configura y retorna un nuevo enrutador
func NewRouter(deps *config.Dependencies) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configurar rutas
	r.Get("/health", handlers.StatusHealth(deps))

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/user", user.HandlerCreateUser(deps))
	})

	return r
}
