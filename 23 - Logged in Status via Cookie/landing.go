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

type User struct {
	FName    string
	LName    string
	Age      string
	HMAC     string
	uuid     *uuid.UUID
	loggedin bool
}

var Useroni User

func landing(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "landing.gohtml", Useroni)
}
