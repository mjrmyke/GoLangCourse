package main

import (
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "github.com/nu7hatch/gouuid"
	// "log"
	"net/http"
	// "strings"
	// "github.com/nu7hatch/gouuid"
	// "text/template"
)

func failedlogin(res http.ResponseWriter, req *http.Request) {
	templates.Execute(res, "failedpage.gohtml")
}
