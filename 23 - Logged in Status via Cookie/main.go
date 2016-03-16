package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"net/http"
	// "strings"
	"text/template"
)

var templates *template.Template
var cook *http.Cookie

func init() {
	templates, _ = template.ParseGlob("templates/*.gohtml") //can not use error on this lol
}

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/failed/", failedlogin)
	http.HandleFunc("/loggedin/", userhome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//Getcookie will get your cookies
func Getcookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	cook, error := req.Cookie("session-fino")
	if error != nil {
		fmt.Println("error: ", error)
		if req.Method == "POST" {
			id, _ := uuid.NewV4()
			cook = givecookie(res, req, req.FormValue("age"), req.FormValue("FName"), id, "True")
			fmt.Println(req.FormValue("FName") + "Asdasd")
			http.SetCookie(res, cook)
		}
	}

	return cook
}

func givecookie(res http.ResponseWriter, req *http.Request, age string, name string, id *uuid.UUID, logged string) *http.Cookie {

	//put all data in string
	stringdata := name + "|" + age + "|" + id.String() + "|" + makehmac(name) + "|" + logged

	//string to json
	data, error := json.Marshal(stringdata)
	if error != nil {
		log.Fatalln(error)
	}

	//json to base64
	b64 := base64.URLEncoding.EncodeToString(data)

	//create the cookie
	cook = &http.Cookie{
		Name:  "session-fino",
		Value: b64 + stringdata,
		// Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(res, cook)

	return cook
}

func makehmac(data string) string {
	x := hmac.New(sha256.New, []byte("Morton's Salt"))
	io.WriteString(x, data)
	return fmt.Sprintf("%x", x.Sum(nil))
}
