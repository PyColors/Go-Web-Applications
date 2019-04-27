package main

import "net/http"

func main() {
	// Fewer lines but does not have functionalitiy as `Demo: http.Handle`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8000", nil)
}
