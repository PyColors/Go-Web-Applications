package webapp

import (
	"net/http"
)

func main() {
	// FileServer method from the http package
	// He takes a root file system
	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
}
