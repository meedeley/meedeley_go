package main

import (
	"log"

	"github.com/meedeley/go-launch-starter-code/internal/configs"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http"
)

func main() {

	logger := log.Default()

	app := configs.RunApp{
		Http: http.Http,
		Cors: "*",
		Log:  logger,
	}

	app.Start()
}
