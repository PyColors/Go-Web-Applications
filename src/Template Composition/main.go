package webapp

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Set template variable to populateTemplate() bellow
	template := populateTemplate()
	// Custom Handle
	// Handle any function into the application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Determine the name of the template
		requestedFile := r.URL.Path[1:]

		// He return a map with the name of the request file
		t := template[(requestedFile + ".html")]

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
// Return map[string]*template : pull and clone a layout for each individual layout that we have
// For every template
func populateTemplate() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "template"

	// Load the template by PareFiles
	layout := template.Must(template.PareFiles(basePath + "/_layout.html"))

	// Load template that layout gonna use by ParseFiles in the layout pull header and footer template
	template.Must(layout.ParseFiles(basePath+ "/_header.html", basePath+ "/_footer.html"))

	// Load the actual template
	// All the content templates will be defined inside content directory
	// Open with the `os` commend from os package
	dir, err := os.Open(basePath + "/content")

	// Error check
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	// Read all of content by `dir` command
	fis, err := dir.Readdir(-1)

	// Error check
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	// Loop all files
	for _, fi := range fis {

		// Open the file where is pointing to
		f, err := os.Open(basePath + "/content/" + fi.Name())

		// Error check
		if err != nil {
			panic("Failed to open templates'" + fi.Name() + "'")
		}

		// Read the content
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to open content from file '" + fi.Name() + "'")
		}

		// Close file
		f.Close()

		// Create the actual template itself
		// Clone method on the layout template
		// He takes the layout template + all children and clone that into `tmpl` object

		tmpl := template.Must(layout.Clone())

		// `tmpl` object he's ready
		// Parse the content been read to that file
		_, err = tmpl.Parse(string(content))

		// Error check as Parse a template might fails
		if err != nil {
			panic("Failed to parse content of '" + fi.Name() + "' as template")
		}

		// Add that template to result map
		result[fi.Name()] = tmpl
	}
	return result
}
