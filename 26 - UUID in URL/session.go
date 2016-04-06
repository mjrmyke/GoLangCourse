package mem

import (
	"errors"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func getID(res http.ResponseWriter, req *http.Request) (string, error) {
	ctx := appengine.NewContext(req)
	var id, origin string
	var cookie *http.Cookie
	origin = "cookerino"
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		origin = "URL"
		id := req.FormValue("id")
		if id == "" {
			origin = "Usercreated via logout"
			log.Infof(ctx, "ID created by: %s", origin)
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return id, errors.New("error logged out manually")
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	log.Infof(ctx, "ID CAME FROM %s", origin)
	return id, nil
}
