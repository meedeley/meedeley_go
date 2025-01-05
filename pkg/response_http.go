package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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

func Response401(w http.ResponseWriter) {
	err := Response{
		Status:  http.StatusNotFound,
		Message: "no found",
	}

	ResponseJSON(w, http.StatusNotFound, err, nil)
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data any, headers http.Header) {

	json, err := json.Marshal(data)

	if err != nil {
		log.Println(err.Error())
		path := "pkg/log/error.log"
		file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(Response{
		Status:  int(statusCode)
		Message: string(json),
	})
}
