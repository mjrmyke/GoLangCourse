package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		if req.Method == "POST" {
			file, _, err := req.FormFile("formy-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer file.Close()
			source := io.LimitReader(file, 400)
			dest, err := os.Create(filepath.Join(".", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dest.Close()
			io.Copy(dest, source)
		}

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<Form method = "POST" enctype="multipart/form-data">
      File Upload:
      <input type="file" name="formy-file">
      <input type="submit">
      </form>
      <br>
      <br>
      `)

	})
	http.ListenAndServe(":8080", nil)
}
