// web.go
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Structure to hold the Page
type Page struct {
	Title string
	Body  []byte
}

// look in to to Responseawriter and hhtp.request why one is pointer and other is value.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //r.URL.Path is the path component of the request URL
}

// load the file
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// hadle the page request if it exist
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err == nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	} else {
		fmt.Fprint(w, "<h1>%s</h1>", title+" page does not exist")
	}
}

// handle for the editing
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}
	htmlTemp, _ := template.ParseFiles("edit.html")
	htmlTemp.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler) //handle all requests to the web root ("/") with handler
	http.HandleFunc("/edit/", editHandler) // handle the edit request.
	http.ListenAndServe(":9090", nil)      //function will block until program is terminated.
}

// Next : https://golang.org/doc/articles/wiki/# Using net/http to serve wiki pages
