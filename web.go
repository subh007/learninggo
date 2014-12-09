// web.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //r.URL.Path is the path component of the request URL
}

func main() {
	http.HandleFunc("/", handler)     //handle all requests to the web root ("/") with handler
	http.ListenAndServe(":9090", nil) //function will block until program is terminated.
}

// Next : https://golang.org/doc/articles/wiki/# Using net/http to serve wiki pages
