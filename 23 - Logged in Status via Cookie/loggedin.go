package main

import (
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "github.com/nu7hatch/gouuid"
	// "log"
	"net/http"
	// "strings"
	// "text/template"
)

func userhome(res http.ResponseWriter, req *http.Request) {
	templates.Execute(res, "loggedin.gohtml")
}
