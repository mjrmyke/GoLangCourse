package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"os"
	"text/template"
)

type User struct {
	FName string
	LName string
	age   string
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {

	useroni := User{}

	templ, error := template.ParseFiles("tpl.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}

	error = templ.Execute(os.Stdout, nil)
	if error != nil {
		log.Fatalln(error)
	}

	cookie, err := req.Cookie("session-fino")

	useroni.FName = req.FormValue("FName")
	useroni.LName = req.FormValue("LName")
	useroni.age = req.FormValue("age")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "_" + useroni.FName + "_" + useroni.LName,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

	templ.ExecuteTemplate(res, "tpl.gohtml", useroni)

}
