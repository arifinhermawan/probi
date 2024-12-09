package router

import (
	"context"
	"net/http"

	blog "github.com/arifinhermawan/blib/log"
)

func routeMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctx := context.WithValue(ctx, blog.ContextKey("url"), r.URL.Path)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
