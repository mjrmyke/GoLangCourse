package main

import (
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	//
	// "log"
	"net/http"
	// "strings"
	"text/template"
)

var templates *template.Template

func init() {
	templates, _ = template.ParseGlob("templates/*.html") //can not use error on this lol
}

func main() {

	http.HandleFunc("/", landing)
	http.HandleFunc("/failed/", failedlogin)
	http.HandleFunc("/loggedin/", userhome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func failedlogin(res http.ResponseWriter, req *http.Request) {
	templates.Execute(res, "failedpage.gohtml")
}
