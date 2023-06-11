package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"test-open-api/internal/api"
	"test-open-api/internal/monitoring"
)

func New(server api.ServerInterface) http.Handler {
	router := chi.NewRouter()

	router.Use(monitoring.ContextMiddleware)
	router.Use(PanicRecoveryMiddleware)

	router.Mount("/swagger", swaggerUIRouter())
	return api.HandlerFromMux(server, router)
}
