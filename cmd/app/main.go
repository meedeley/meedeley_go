package app

import (
	"github.com/meedeley/go-launch-starter-code/internal/delivery/http"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

func main() {

	pkg.ZeroLog()
	http.Http()

}
