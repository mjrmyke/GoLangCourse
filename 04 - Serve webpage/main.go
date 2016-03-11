//Myke Walker

package main

import (
	"html/template"
	"log"
	"net/http"
)

func serve(res http.ResponseWriter, req *http.Request) {
	templ := template.New("index.html")
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	templ.Execute(res, nil)
}

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}
