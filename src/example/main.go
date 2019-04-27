package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

// Method takes two argument ResponseWriter and a pointer http.Request
func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "Hello"})
	http.ListenAndServe(":8000", nil)
}
