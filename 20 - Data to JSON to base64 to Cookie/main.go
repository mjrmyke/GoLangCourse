package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	// "os"
	"text/template"
)

type User struct {
	FName string
	LName string
	Age   string
	uuid  *uuid.UUID
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

	//make uuid
	id, _ := uuid.NewV4()

	cookie, err := req.Cookie("session-fino")

	//get data from form
	useroni.FName = req.FormValue("FName")
	useroni.LName = req.FormValue("LName")
	useroni.Age = req.FormValue("age")
	useroni.uuid = id

	//encode data to json
	data, error := json.Marshal(useroni)
	if error != nil {
		log.Fatalln(error)
	}

	//json data to base64
	b64 := base64.URLEncoding.EncodeToString(data)

	if err != nil {
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: b64,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

	templ.ExecuteTemplate(res, "tpl.gohtml", useroni)

}
