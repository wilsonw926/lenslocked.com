package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<html>To get in touch, email me at <a href=\"mailto:test@gmail.com\">test@gmail.com</a>.</html>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
