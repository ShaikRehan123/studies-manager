package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}, message string, status int) {
	response := Response{
		Success: status >= 200 && status < 300,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		SendErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SendErrorResponse(w http.ResponseWriter, message string, status int) {
	response := ErrorResponse{
		Success: false,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SendSuccess(w http.ResponseWriter, data interface{}, message string) {
	SendJSONResponse(w, data, message, http.StatusOK)
}

func SendCreated(w http.ResponseWriter, data interface{}, message string) {
	SendJSONResponse(w, data, message, http.StatusCreated)
}

func SendBadRequest(w http.ResponseWriter, message string) {
	SendErrorResponse(w, message, http.StatusBadRequest)
}

func SendNotFound(w http.ResponseWriter, message string) {
	SendErrorResponse(w, message, http.StatusNotFound)
}

func SendInternalError(w http.ResponseWriter, message string) {
	SendErrorResponse(w, message, http.StatusInternalServerError)
}
