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
func populateTemplate() map[string]*template.Template {


	return result
}
