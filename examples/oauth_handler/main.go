package main

import (
	"net/http"
	"os"

	"github.com/heroku/authenticater"
	"github.com/kr/secureheader"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	behindGoogleOAuth := &authenticater.OAuthHandler{
		RequireDomain: os.Getenv("REQUIRE_DOMAIN"),
		Key:           os.Getenv("KEY"),
		ClientID:      os.Getenv("CLIENT_ID"),
		ClientSecret:  os.Getenv("CLIENT_SECRET"),
		Handler:       mux,
	}

	http.Handle("/", behindGoogleOAuth)
	http.ListenAndServe(":"+os.Getenv("PORT"), secureheader.DefaultConfig)
}
