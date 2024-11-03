package router

import (
	"context"
	"net/http"

	blog "github.com/arifinhermawan/blib/log"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func routeMiddleware(app *newrelic.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(r.URL.Path)
		defer txn.End()

		txn.SetWebRequestHTTP(r)
		w = txn.SetWebResponse(w)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		customWriter := &statusTrackingResponseWriter{ResponseWriter: w}

		ctx := internalContext.DefaultContext()
		ctx = context.WithValue(ctx, blog.ContextKey("url"), r.URL.Path)
		r = r.WithContext(ctx)

		next.ServeHTTP(customWriter, r)
	})
}

type statusTrackingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusTrackingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
