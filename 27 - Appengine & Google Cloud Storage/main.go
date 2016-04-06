package mykecloud

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

const gcsBucket = "testenv-1273.appspot.com"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/golden", retriever)
}

func handler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
		<h1>UPLOAD</h1>
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {

		mpf, hdr, err := req.FormFile("dahui")
		if err != nil {
			log.Errorf(ctx, "ERROR handler req.FormFile: ", err)
			http.Error(res, "We were unable to upload your file\n", http.StatusInternalServerError)
			return
		}
		defer mpf.Close()

		fname, err := uploadFile(req, mpf, hdr)
		if err != nil {
			log.Errorf(ctx, "ERROR handler uploadFile: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		fnames, err := putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		html += `<h1>Files</h1>`
		for k, _ := range fnames {
			html += `<h3><a href="/golden?object=` + k + `">` + k + `</a></h3>`
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}

func retriever(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	objectName := req.FormValue("object")
	rdr, err := getFile(ctx, objectName)
	if err != nil {
		log.Errorf(ctx, "ERROR golden getFile: ", err)
		http.Error(res, "We were unable to get the file"+objectName+"\n"+err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer rdr.Close()
	io.Copy(res, rdr)
}

func uploadFile(req *http.Request, mpf multipart.File, hdr *multipart.FileHeader) (string, error) {

	ext, err := fileFilter(req, hdr)
	if err != nil {
		return "", err
	}
	name := getSha(mpf) + `.` + ext
	mpf.Seek(0, 0)

	ctx := appengine.NewContext(req)
	return name, putFile(ctx, name, mpf)
}

func fileFilter(req *http.Request, hdr *multipart.FileHeader) (string, error) {

	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "FILE EXTENSION: %s", ext)

	switch ext {
	case "jpg", "jpeg", "txt", "md":
		return ext, nil
	}
	return ext, fmt.Errorf("We do not allow files of type %s. We only allow jpg, jpeg, txt, md extensions.", ext)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}

func getFile(ctx context.Context, name string) (io.ReadCloser, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(gcsBucket).Object(name).NewReader(ctx)
}

func putCookie(res http.ResponseWriter, req *http.Request, fname string) (map[string]bool, error) {
	mss := make(map[string]bool)
	cookie, _ := req.Cookie("file-names")
	if cookie != nil {
		bs, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler base64.URLEncoding.DecodeString: %s", err)
		}
		err = json.Unmarshal(bs, &mss)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler json.Unmarshal: %s", err)
		}
	}

	mss[fname] = true
	bs, err := json.Marshal(mss)
	if err != nil {
		return mss, fmt.Errorf("ERROR putCookie json.Marshal: ", err)
	}
	b64 := base64.URLEncoding.EncodeToString(bs)

	// FYI
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "COOKIE JSON: %s", string(bs))

	http.SetCookie(res, &http.Cookie{
		Name:  "file-names",
		Value: b64,
	})
	return mss, nil
}
