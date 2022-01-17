package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Hello， this is goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "abouttttt")
	} else {
		fmt.Fprintf(w, "<h1>11111NOT FOUND</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
