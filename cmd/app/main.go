package main

import (
	"github.com/meedeley/go-launch-starter-code/internal/deliveries/http"
)

func main() {
	run := http.Http()

	run.Listen(":3000")
}
