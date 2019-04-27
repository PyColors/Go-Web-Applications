package webapp

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Looking into the request Object, see what file was requested
		// and then, translate that over to the file system

		// `os.Open()` from os package
		f, err := os.Open("public" + r.URL.Path)

		// CHek if file was open successfully
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		// if everything did succeed
		// CLose file off
		defer f.Close()

		// New variable `contentType`
		// Looking the suffix of the incoming URL
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"

		}
		// Once the contentType has been defined
		// Add the Herder to the response
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
