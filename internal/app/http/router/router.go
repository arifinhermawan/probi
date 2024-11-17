package router

import (
	"context"
	"log"
	"net/http"

	"github.com/arifinhermawan/probi/internal/app/http/server"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func HandleRequest(ctx context.Context, lib *lib.Lib, handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(lib, handlers, router)
	handlePutRequest(lib, handlers, router)
	handlePostRequest(lib, handlers, router)

	log.Println("SERVING AT PORT :8080")
	log.Fatal(http.ListenAndServe(newrelic.WrapListen(":8080"), routeMiddleware(ctx, router)))
}

func handleGetRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	// Reminder endpoints
	router.HandleFunc("/reminder", lib.AuthMiddleware(handlers.Reminder.GetUserActiveReminderHandler)).Methods("GET")

	// User endpoints
	router.HandleFunc("/user/{user_id}", lib.AuthMiddleware(handlers.User.GetUserDetailsHandler)).Methods("GET")
}

func handlePutRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	// Reminder endpoints
	router.HandleFunc("/reminder", lib.AuthMiddleware(handlers.Reminder.UpdateReminderHandler)).Methods("PUT")

	// User endpoints
	router.HandleFunc("/user", lib.AuthMiddleware(handlers.User.UpdateUserDetailsHandler)).Methods("PUT")
}

func handlePostRequest(lib *lib.Lib, handlers *server.Handlers, router *mux.Router) {
	// Auth endpoints
	router.HandleFunc("/auth/login", handlers.Auth.LogInHandler).Methods("POST")
	router.HandleFunc("/auth/logout", lib.AuthMiddleware(handlers.Auth.LogOutHandler)).Methods("POST")

	// Reminder endpoints
	router.HandleFunc("/reminder", lib.AuthMiddleware(handlers.Reminder.CreateReminderHandler)).Methods("POST")

	// User endpoints
	router.HandleFunc("/user", handlers.User.CreateUserHandler).Methods("POST")
}
