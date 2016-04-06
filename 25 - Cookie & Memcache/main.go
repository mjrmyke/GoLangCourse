package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"

	"github.com/nu7hatch/gouuid"
)

type model struct {
	Name string
	ID   *uuid.UUID
}

func init() {
	// templates, _ := template.ParseGlob("*.gohtml")
	http.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, error := req.Cookie("cook-o-matic")
	if error != nil {
		id := cookienidcreation(res, req)
		ctx := appengine.NewContext(req)
		item1 := memcache.Item{
			Key:   id,
			Value: []byte("Myke"),
		}
		memcache.Set(ctx, &item1)
	}

}

func cookienidcreation(res http.ResponseWriter, req *http.Request) string {

	id, _ := uuid.NewV4()

	//create the cookie
	cook := &http.Cookie{
		Name:  "cook-o-matic",
		Value: id.String(),
		// Secure: true,
		HttpOnly: true,
	}

	http.SetCookie(res, cook)
	return id.String()
}
