package routers

import (
	"go.uber.org/zap"
	"net/http"
)

var logger, _ = zap.NewProduction()

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error("ERROR LOGGER")
				}
			}()
			next.ServeHTTP(w, r)
		})
}
