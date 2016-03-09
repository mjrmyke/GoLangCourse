package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {

	templ, error := template.ParseFiles("tpl.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}

	error = templ.Execute(os.Stdout, nil)
	if error != nil {
		log.Fatalln(error)
	}

	cookie, err := req.Cookie("session-fino")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

	templ.Execute(res, "tpl.gohtml")

}
