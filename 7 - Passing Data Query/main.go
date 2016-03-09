package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		value := req.URL.Query().Get("q")
		io.WriteString(res, "search is: "+value)
	})
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/?q=adsda loads -> search is: adsda
