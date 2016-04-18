package mykecloud

import (
	"io"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

const gcsBucket = "testenv-1273.appspot.com"

//create  struct to store relevant gcs and gae references
type googctx struct {
	ctx      context.Context
	response http.ResponseWriter
	bucket   *storage.BucketHandle
	client   *storage.Client
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {
	//create context needed for app engine and gcs
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "error making new client")
	}
	defer client.Close()

	//place relevent info into ADT
	googappengine := &googctx{
		ctx:      ctx,
		response: res,
		bucket:   client.Bucket(gcsBucket),
		client:   client,
	}

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="uploader">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {

		mpf, hdr, err := req.FormFile("uploader")
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
			html += `<h3>` + k + `</h3>`
		}
	}
	//bobs photos
	user := "bob"
	html += "BOBS PHOTOS:"
	query := &storage.Query{
		Prefix: user + "/",
	}

	objs, err := googappengine.bucket.List(googappengine.ctx, query)
	if err != nil {
		log.Errorf(googappengine.ctx, "ERROR in query_delimiter")
		return
	}

	for _, obj := range objs.Results {
		html += `<img src=https://storage.googleapis.com/testenv-1273.appspot.com/` + obj.Name + `>` + "\n"
	}
	html += "\n\n\n\n"

	//staceys photos
	user = "stacey"
	html += "staceys PHOTOS:\n"
	query = &storage.Query{
		Prefix: user + "/",
	}

	objs, err = googappengine.bucket.List(googappengine.ctx, query)
	if err != nil {
		log.Errorf(googappengine.ctx, "ERROR in query_delimiter")
		return
	}

	for _, obj := range objs.Results {
		html += `<img src=https://storage.googleapis.com/testenv-1273.appspot.com/` + obj.Name + `>` + "\n"
	}

	html += "\n\n\n\n"

	//james' photos
	user = "james"
	html += "james' PHOTOS: \n"
	query = &storage.Query{
		Prefix: user + "/",
	}

	objs, err = googappengine.bucket.List(googappengine.ctx, query)
	if err != nil {
		log.Errorf(googappengine.ctx, "ERROR in query_delimiter")
		return
	}

	for _, obj := range objs.Results {
		html += `<img src=https://storage.googleapis.com/testenv-1273.appspot.com/` + obj.Name + `>` + "\n"
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}
