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
	"text/template"

	"github.com/nu7hatch/gouuid"
)

type User struct {
	FName string
	LName string
	Age   string
	HMAC  string
	uuid  *uuid.UUID
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/failed/", bar)
	http.ListenAndServe(":8080", nil)

}

func makehmac(data string) string {
	x := hmac.New(sha256.New, []byte("salt"))
	io.WriteString(x, data)
	return fmt.Sprintf("%x", x.Sum(nil))
}

func bar(res http.ResponseWriter, req *http.Request) {

	templ2, error := template.ParseFiles("tpl2.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}
	templ2.Execute(res, "tpl2.gohtml")

}

func foo(res http.ResponseWriter, req *http.Request) {

	templ, error := template.ParseFiles("tpl.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}

	cookie, err := req.Cookie("session-fino")

	useroni := User{}

	//if no cookie
	if err != nil {

		//make uuid
		id, _ := uuid.NewV4()
		//get data from form
		useroni.FName = req.FormValue("FName")
		useroni.LName = req.FormValue("LName")
		useroni.Age = req.FormValue("Age")
		useroni.HMAC = makehmac(useroni.FName)
		useroni.uuid = id

		//encode data to json
		stringdata := useroni.FName + "|" + useroni.LName + "|" + useroni.Age + "|" + useroni.uuid.String() + "|" + useroni.HMAC
		data, error := json.Marshal(stringdata)
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

		//no cookie
	} else {
		var rcvduseroni string
		usercookiedata := cookie.Value
		decodeduserdata, err := base64.URLEncoding.DecodeString(usercookiedata)
		if err != nil {
			log.Println("Error decoding base64", err)
		}
		err = json.Unmarshal(decodeduserdata, &rcvduseroni)
		if err != nil {
			fmt.Println("error unmarshalling: ", err)
		}
		splitstrings := strings.Split(rcvduseroni, "|")
		rcvdfname := splitstrings[0]
		rcvdlname := splitstrings[1]
		rcvdage := splitstrings[2]
		rcvdhmac := splitstrings[4]
		codeCheck := makehmac(rcvdfname)
		rcvdhmac = rcvdhmac
		// rcvdhmac = rcvdhmac + "failing  verification lol"

		if codeCheck != rcvdhmac {
			log.Println("MISMATCHED UUID HMAC")
			log.Println(rcvdhmac)
			log.Println(codeCheck)
			http.Redirect(res, req, "/failed/", http.StatusMovedPermanently)
		}
		log.Println(rcvdfname + rcvdlname + rcvdage)

	}
	templ.ExecuteTemplate(res, "tpl.gohtml", useroni)

}
