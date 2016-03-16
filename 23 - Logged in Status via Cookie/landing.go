package main

import (
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "github.com/nu7hatch/gouuid"
	// "log"
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
	loggedin string
}

//Useroni Declared as empty user
var Useroni User

func landing(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "landing.gohtml", Useroni)
	Getcookie(res, req)
}
