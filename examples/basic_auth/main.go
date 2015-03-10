package main

import (
	"net/http"

	"github.com/heroku/authenticater"
)

func main() {
	auth := authenticater.NewBasicAuth()
	auth.AddPrinciple("foo", "bar")
	http.HandlerFunc("/", authenticater.WrapAuth(auth, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}))
	http.ListenAndServe(":8080", nil)
}
