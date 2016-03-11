package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This should be in HTTPS! Shiny! -- if you get a certificate lol\n"))
}

func main() {
	http.HandleFunc("/", home)
	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", 301))
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
