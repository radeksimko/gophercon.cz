package app

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", redirectToGithub)
	http.HandleFunc("/robots.txt", plainText)
}

func redirectToGithub(w http.ResponseWriter, r *http.Request) {
	h := http.RedirectHandler("https://github.com/golang/go/wiki/Conferences", 302)
	h.ServeHTTP(w, r)
}

func plainText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
}
