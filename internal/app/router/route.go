package utils

import (
	"log"
	"net/http"

	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/gorilla/mux"
)

func HandleRequest(handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(handlers, router)
	handlePatchRequest(handlers, router)
	handlePostRequest(handlers, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleGetRequest(handlers *server.Handlers, router *mux.Router) {
}

func handlePatchRequest(handlers *server.Handlers, router *mux.Router) {
}

func handlePostRequest(handlers *server.Handlers, router *mux.Router) {
}
