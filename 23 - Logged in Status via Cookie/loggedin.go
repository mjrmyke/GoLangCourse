package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	// "text/template"
)

func userhome(res http.ResponseWriter, req *http.Request) {

	_, error := req.Cookie("session-fino")
	if error != nil || Useroni.Loggedin != "True" {
		log.Println("Error detecting cookie in logged in page"+Useroni.Loggedin, error)
		http.Redirect(res, req, "/failed/", 302)
	} else {
		value := req.URL.Query().Get("logout")
		if value == "logout" {
			http.Redirect(res, req, "/loggedout", 302)
			Useroni.Loggedin = "False"
			givecookie(res, req, Useroni.Age, Useroni.FName, Useroni.Uuid, "False")
			Getcookie(res, req)

		}
		if req.Method == "POST" {
			file, _, err := req.FormFile("formy-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer file.Close()
			source := io.LimitReader(file, 400)
			dest, err := os.Create(filepath.Join(".", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dest.Close()
			io.Copy(dest, source)
		}
		templates.ExecuteTemplate(res, "loggedin.gohtml", Useroni)
	} //end else
} //end userhome
