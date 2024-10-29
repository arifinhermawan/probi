package handler

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, code int, result interface{}, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	response := defaultResponse{
		Code:    code,
		Message: message,
	}
	if result != nil {
		response.Response = result
	}
	if err != nil {
		response.Error = err.Error()
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
