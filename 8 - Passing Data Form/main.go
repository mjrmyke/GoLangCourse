package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		value := req.FormValue("formy")

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<Form method = "POST">
      Name:
      <input type="text" name="formy">
      <input type="submit">
      </form>
      <br>
      <br>
      `+value)

	})
	http.ListenAndServe(":8080", nil)
}
