package trace

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"go.uber.org/zap"
)

const key = "logger"

func NewTraceMiddleware(logger *zap.Logger) mux.MiddlewareFunc{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceparent := r.Header.Get("traceparent")
			// l := logger.
			l := logger.With(zap.String("traceparent",traceparent))
			next.ServeHTTP(w,r.WithContext(context.WithValue(r.Context(), key, l)))
		})
	}
}

func UnWrap(r *http.Request) *zap.Logger {
	val := r.Context().Value(key)
	if logger, ok := val.(*zap.Logger); ok {
		return logger
	}
	return zap.NewExample()
}