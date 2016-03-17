package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/nu7hatch/gouuid"
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
	http.HandleFunc("/login/", login)
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
			cook = givecookie(res, req, req.FormValue("Age"), req.FormValue("FName"), id, "True")
			fmt.Println(req.FormValue("FName"))
			http.SetCookie(res, cook)

			//Places new cookie data into user struct "Useroni"
			var rcvduseroni string

			usercookiedata := cook.Value
			decodeduserdata, err := base64.URLEncoding.DecodeString(usercookiedata)
			if err != nil {
				log.Println("Error decoding base64", err)
			}
			err = json.Unmarshal(decodeduserdata, &rcvduseroni)
			if err != nil {
				fmt.Println("error unmarshalling: ", err)
			}

			splitstrings := strings.Split(rcvduseroni, "|")

			Useroni.FName = splitstrings[0]
			Useroni.Age = splitstrings[1]

			parseduuid, err := uuid.ParseHex(splitstrings[2])
			if err != nil {
				fmt.Println("error parsing uuid from cookie: ", err)
			}

			Useroni.Uuid = parseduuid
			Useroni.HMAC = splitstrings[3]
			Useroni.Loggedin = splitstrings[4]
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
		Value: b64,
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
