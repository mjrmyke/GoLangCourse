package main

import (
	"io"
	"net/http"
	"strings"
)

func handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	name := strings.Split(req.URL.Path, "/")
	io.WriteString(res, name[1])
}

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
