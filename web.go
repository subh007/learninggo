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

// saving the page
func (p *Page) save() { // difference between func (p *Page) and func (p Page)
	filename := p.Title + ".txt"
	ioutil.WriteFile(filename, p.Body, 0600)
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
		// page exist
		htmlTemp, _ := template.ParseFiles("view.html")
		htmlTemp.Execute(w, p)
	} else {
		// page does not exist
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
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

// saving the page
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	p := Page{Title: title, Body: []byte(body)}
	p.save()

	http.Redirect(w, r, "/view/"+title, http.StatusFound) // what is statusfound
}

func main() {
	http.HandleFunc("/view/", viewHandler) //handle all requests to the web root ("/") with handler
	http.HandleFunc("/edit/", editHandler) // handle the edit request.
	http.HandleFunc("/save/", saveHandler) // handle saving the page.
	http.ListenAndServe(":9090", nil)      //function will block until program is terminated.
}

// Next : https://golang.org/doc/articles/wiki/# Saving Pages
