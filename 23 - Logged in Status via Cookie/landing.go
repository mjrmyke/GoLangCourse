package main

import (
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	// "strings"
	"github.com/nu7hatch/gouuid"
	// "text/template"
)

//User type
type User struct {
	FName    string
	Age      string
	Uuid     *uuid.UUID
	HMAC     string
	Loggedin string
}

//Useroni Declared as empty user
var Useroni User

func landing(res http.ResponseWriter, req *http.Request) {
	Getcookie(res, req)
	templates.ExecuteTemplate(res, "landing.gohtml", Useroni)
}

func login(res http.ResponseWriter, req *http.Request) {
	Getcookie(res, req)
	givecookie(res, req, Useroni.Age, Useroni.FName, Useroni.Uuid, "True")
	Getcookie(res, req)
	Useroni.Loggedin = "True"
	log.Println(cook.Value)
	// http.Redirect(res, req, "/loggedin/", 302)
	templates.ExecuteTemplate(res, "login.gohtml", Useroni)

}
