package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http"
)

func main() {
	run := http.Http()

	run.Listen(":3000")
}
