package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func makehmac(data string) string {
	x := hmac.New(sha256.New, []byte("key"))
	io.WriteString(x, data)
	return fmt.Sprintf("%x", x.Sum(nil))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/authenticate", authenticate)
	http.ListenAndServe(":8080", nil)

}

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	cook, err := req.Cookie("session-fino")

	if err != nil {
		cook = &http.Cookie{
			Name:  "session-fino",
			Value: "",
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("formtext") != "" {
		formtext := req.FormValue("formtext")
		cook.Value = formtext + `|` + makehmac(formtext)
	}

	http.SetCookie(res, cook)
	io.WriteString(res, `<!DOCTYPE html>
  	<html>
  	  <body>
  	    <form method="POST">
  	    `+cook.Value+`
  	      <input type="text" name="formtext">
  	      <input type="submit">
  	    </form>
  	    <a href="/authenticate">Validate This HMAC</a>
  	  </body>
  	</html>`)

}

func authenticate(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-fino")
	if err != nil {
		http.Redirect(res, req, "/", 303)
		return
	}

	if cookie.Value == "" {
		http.Redirect(res, req, "/", 303)
		return
	}

	data := strings.Split(cookie.Value, "|")
	formtext := data[0]
	codeRcvd := data[1]
	codeCheck := makehmac(formtext)
	//codeCheck := makehmac(formtext + "s")
	//if uncommented, this line will invalidate the HMAC

	if codeRcvd != codeCheck {
		log.Println("Invalidated HMAC CODE - ACCESS DENIED!!")
		log.Println(codeRcvd)
		log.Println(codeCheck)
		http.Redirect(res, req, "/", 303)
		return
	}

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeRcvd+` - USER COOKIE CODE </h1>
	  	<h1>`+codeCheck+` - EXPECTED COOKIE CODE </h1>
	  </body>
	</html>`)
}
