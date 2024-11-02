package utils

import (
	"context"
	"log"
	"net/http"

	blog "github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/arifinhermawan/probi/internal/lib"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
	"github.com/gorilla/mux"
)

func HandleRequest(lib *lib.Lib, handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(lib, handlers, router)
	handlePatchRequest(lib, handlers, router)
	handlePostRequest(lib, handlers, router)

	log.Println("SERVING AT PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", routeMiddleware(router)))
}

func routeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctx := internalContext.DefaultContext()
		ctx = context.WithValue(ctx, blog.ContextKey("url"), r.URL.Path)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func handleGetRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	router.HandleFunc("/user/{user_id}", lib.AuthMiddleware(handlers.User.GetUserDetailsHandler)).Methods("GET")
}

func handlePatchRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
}

func handlePostRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	// Auth endpoints
	router.HandleFunc("/auth/login", handlers.Auth.LogInHandler).Methods("POST")
	router.HandleFunc("/auth/logout", lib.AuthMiddleware(handlers.Auth.LogOutHandler)).Methods("POST")

	// User endpoints
	router.HandleFunc("/user", handlers.User.CreateUserHandler).Methods("POST")
}
