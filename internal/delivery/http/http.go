package http

import "net/http"

func SetUpDB() {

}

func Http() {

	SetUpDB()

	http.HandleFunc("/")
}
