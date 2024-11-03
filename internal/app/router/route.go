package router

import (
	"log"
	"net/http"

	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func HandleRequest(app *newrelic.Application, lib *lib.Lib, handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(lib, handlers, router)
	handlePutRequest(lib, handlers, router)
	handlePostRequest(lib, handlers, router)

	log.Println("SERVING AT PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", routeMiddleware(app, router)))
}

func handleGetRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	router.HandleFunc("/user/{user_id}", lib.AuthMiddleware(handlers.User.GetUserDetailsHandler)).Methods("GET")
}

func handlePutRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	router.HandleFunc("/user", lib.AuthMiddleware(handlers.User.UpdateUserDetailsHandler)).Methods("PUT")
}

func handlePostRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	// Auth endpoints
	router.HandleFunc("/auth/login", handlers.Auth.LogInHandler).Methods("POST")
	router.HandleFunc("/auth/logout", lib.AuthMiddleware(handlers.Auth.LogOutHandler)).Methods("POST")

	// User endpoints
	router.HandleFunc("/user", handlers.User.CreateUserHandler).Methods("POST")
}
