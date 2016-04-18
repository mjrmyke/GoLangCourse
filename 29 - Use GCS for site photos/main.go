package mykecloud

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"io/ioutil"
	"net/http"
)

//make constant for easy bucket reference
const bucket = "testenv-1273.appspot.com"

var tpl *template.Template

//create  struct to store relevant gcs and gae references
type googctx struct {
	ctx      context.Context
	response http.ResponseWriter
	bucket   *storage.BucketHandle
	client   *storage.Client
}

func init() {
	//html files, and strips prefixes for img and css
	tpl, _ = template.ParseGlob("*.html")
	http.HandleFunc("/", main)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img"))))
}

func main(response http.ResponseWriter, request *http.Request) {
	//restricts access to anything other than landing page
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	//create context needed for app engine and gcs
	ctx := appengine.NewContext(request)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "error making new client")
	}
	defer client.Close()

	//place relevent info into ADT
	googappengine := &googctx{
		ctx:      ctx,
		response: response,
		bucket:   client.Bucket(bucket),
		client:   client,
	}

	//call function to upload all images
	googappengine.uploadFiles()

	//string slice photos
	var listedphotos []string

	//make query with requested parameters
	query := &storage.Query{
		Prefix:    "photos/",
		Delimiter: "/",
	}

	//get objects from query
	objs, err := googappengine.bucket.List(googappengine.ctx, query)
	if err != nil {
		log.Errorf(googappengine.ctx, "ERROR in query_delimiter")
		return
	}

	//place links in string
	for _, obj := range objs.Results {
		listedphotos = append(listedphotos, obj.MediaLink)
	}

	//Serves pages with string slice required to place all images on page
	tpl.ExecuteTemplate(response, "index.html", listedphotos)
}

//function for uploading files
func (googappengine *googctx) uploadFiles() {
	for _, fileName := range []string{
		"photos/0.jpg",
		"photos/2.jpg",
		"photos/3.jpg",
		"photos/3.jpg",
		"photos/4.jpg",
		"photos/5.jpg",
		"photos/6.jpg",
		"photos/7.jpg",
		"photos/8.jpg",
		"photos/10.jpg",
		"photos/11.jpg",
		"photos/12.jpg",
		"photos/13.jpg",
		"photos/14.jpg",
		"photos/15.jpg",
		"photos/16.jpg",
		"photos/17.jpg",
		"photos/18.jpg",
		"photos/19.jpg",
		"photos/20.jpg",
		"photos/21.jpg",
		"photos/22.jpg",
		"photos/23.jpg",
		"photos/24.jpg",
	} {
		googappengine.copyFile(fileName)
	}
}

//Copys files to bucket
func (googappengine *googctx) copyFile(fileName string) {
	writer := googappengine.bucket.Object(fileName).NewWriter(googappengine.ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	writer.ContentType = "image/jpg"

	//read all images
	files, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf(googappengine.ctx, "Error in copyFile: ", err)
		return
	}

	//writes images previously read to gcs
	_, err = writer.Write(files)
	if err != nil {
		log.Errorf(googappengine.ctx, "createFile: unable to write data to bucket")
		return
	}
	err = writer.Close()
	if err != nil {
		log.Errorf(googappengine.ctx, "createFile Close")
		return
	}
}
