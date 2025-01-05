package lib

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, message string, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := GenericResponse{
		Message: message,
		Status:  status,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
