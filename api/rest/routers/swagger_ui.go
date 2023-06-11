package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	"net/http"
	"strconv"
	"test-open-api/internal/settings"
	"time"
)

func swaggerUIRouter() chi.Router {
	const swaggerFile = "swagger.yaml"

	cfg := settings.GetConfig()

	r := chi.NewRouter()

	r.Use(
		middleware.BasicAuth("swagger-ui", map[string]string{
			cfg.SwaggerUI.Login: cfg.SwaggerUI.Password,
		}),
	)

	r.Get("/"+swaggerFile, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerFile)
	})

	ts := strconv.FormatInt(time.Now().Unix(), 10)

	r.Get("/ui/*", httpSwagger.Handler(
		// Добавляем к URL timestamp для предотвращения кеширования файла в
		//браузер
		httpSwagger.URL("/swagger/"+swaggerFile+"?ts"+ts),
	))
	return r
}
