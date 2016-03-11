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

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		cook.Value = email + `|` + makehmac(email)
	}

	http.SetCookie(res, cook)
	io.WriteString(res, `<!DOCTYPE html>
  	<html>
  	  <body>
  	    <form method="POST">
  	    `+cook.Value+`
  	      <input type="text" name="email">
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

	xs := strings.Split(cookie.Value, "|")
	fmt.Println("HERE'S THE SLICE", xs)
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := makehmac(email)
	//codeCheck := makehmac(email + "s")

	if codeRcvd != codeCheck {
		log.Println("HMAC codes didn't match")
		log.Println(codeRcvd)
		log.Println(codeCheck)
		http.Redirect(res, req, "/", 303)
		return
	}

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheck+` - RECALCULATED </h1>
	  </body>
	</html>`)
}
