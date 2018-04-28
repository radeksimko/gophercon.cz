package hello

import (
	"net/http"
)

func init() {
	h := http.RedirectHandler("https://github.com/golang/go/wiki/Conferences", 302)
	http.HandleFunc("/", h.ServeHTTP)
}
