package webapp

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Set template variable to populateTemplate() bellow
	template := populateTemplate()
	// Custom Handle
	// Handle any function into the application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Determine the name of the template
		requestedFile := r.URL.Path[1:]

		// Look up the template
		// `Lookup()` function he takes the name of the template
		t := template.Lookup(requestedFile + ".html")

		// if a template was not find
		// happy path
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			// return the status no found if no template been found 404 form http package
			w.WriteHeader(http.StatusNotFound)
		}
	})

	// handle img, css by intense of the file system with a public root
	http.handle("/img", http.FileServer(http.Dir("public")))
	http.handle("/css", http.FileServer(http.Dir("public")))

	// Listen and serve
	http.ListenAndServe("8080", nil)
}

// Allow populating the template
// Function without any parameter in
func populateTemplate() *template.Template {
	// Name will be a container to load template in
	result := template.New("Templates")

	// Where the template is on the file system
	const basePath = "templates"

	// Parse those template in the context or result template
	// with the must function, he takes and a `Template` and an `error`
	// and return a template if succeed or fails if anything fails
	// `ParseGlob` takes a file pattern with the base path
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}
