package main

import "github.com/meedeley/go-launch-starter-code/internal/delivery/http"

func main() {
	run := http.Http()

	run.Listen(":3000")
}
