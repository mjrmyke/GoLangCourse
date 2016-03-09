package main

import (
	"io"
	"net/http"
)

type num int

func (h num) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {
	var temp num

	mux := http.NewServeMux()
	mux.Handle("/", temp)
	http.ListenAndServe(":8080", mux)
}
