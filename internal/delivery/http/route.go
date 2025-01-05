package http

import (
	"log"
	"net/http"

	"github.com/meedeley/go-launch-starter-code/internal/configs"
)

func SetUpDB() {
	_, err := configs.NewDB()

	if err != nil {
		log.Printf("cannot init database %v", err)
	}
}

func Http() error {

	SetUpDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

}
