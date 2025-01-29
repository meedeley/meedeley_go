package main

import (
	"log"

	"github.com/meedeley/go-launch-starter-code/internal/deliveries/http"
)

func main() {
	app := http.Http()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
