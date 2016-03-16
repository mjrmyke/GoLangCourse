package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"strings"
	// "text/template"
)

func userhome(res http.ResponseWriter, req *http.Request) {
	cook := Getcookie(res, req)
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
	// Useroni.loggedin = splitstrings[4]

	templates.Execute(res, "loggedin.gohtml")
}
