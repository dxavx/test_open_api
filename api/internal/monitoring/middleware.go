package monitoring

import (
	"context"
	"github.com/satori/uuid"
	"net/http"
)

type MiddlewareContext string

const (
	traceID MiddlewareContext = "trace_id"
)

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				traceID,
				uuid.NewV4().String(),
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
