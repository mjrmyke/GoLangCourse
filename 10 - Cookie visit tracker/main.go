package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
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
		// id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: "0",
			// Secure: true,
			HttpOnly: true,
		}

	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	fmt.Println(cookie)

	http.SetCookie(res, cookie)

	templ.ExecuteTemplate(res, "tpl.gohtml", cookie)

}
