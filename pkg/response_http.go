package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Response404(w http.ResponseWriter) {
	err := Response{
		Status:  http.StatusNotFound,
		Message: "Resource not found",
	}

	ResponseJSON(w, http.StatusNotFound, err, nil)
}

func Response409(w http.ResponseWriter) {
	err := Response{
		Status:  http.StatusConflict,
		Message: "Conflict occurred",
	}

	ResponseJSON(w, http.StatusConflict, err, nil)
}

func Response500(w http.ResponseWriter) {
	err := Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	}

	ResponseJSON(w, http.StatusInternalServerError, err, nil)
}

func Response401(w http.ResponseWriter) {
	err := Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	}

	ResponseJSON(w, http.StatusUnauthorized, err, nil)
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data any, headers http.Header) {
	response := Response{
		Status: statusCode,
		Data:   data,
	}

	json, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error serializing JSON: %v", err)
		Response500(w)
		return
	}

	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	w.Write(json)
}
