package utils

import (
	"log"
	"net/http"

	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/gorilla/mux"
)

func HandleRequest(handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(cors)

	handleGetRequest(handlers, router)
	handlePatchRequest(handlers, router)
	handlePostRequest(handlers, router)

	log.Println("SERVING AT PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleGetRequest(handlers *server.Handlers, router *mux.Router) {
}

func handlePatchRequest(handlers *server.Handlers, router *mux.Router) {
}

func handlePostRequest(handlers *server.Handlers, router *mux.Router) {
	// User endpoints
	router.HandleFunc("/user", handlers.User.CreateUserHandler).Methods("POST")

	// Auth endpoints
	router.HandleFunc("/login", handlers.Auth.LogInHandler).Methods("POST")
}
