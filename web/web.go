// web.go
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Structure to hold the Page
type Page struct {
	Title  string
	Body   []byte
	Output []byte
}

// saving the page
func (p *Page) save() { // difference between func (p *Page) and func (p Page)
	filename := p.Title + ".go"
	ioutil.WriteFile(filename, p.Body, 0777)
}

// look in to to Responseawriter and hhtp.request why one is pointer and other is value.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //r.URL.Path is the path component of the request URL
}

// load the file
func loadPage(title string) (*Page, error) {
	filename := title + ".go"
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
	http.Redirect(w, r, "/exec/"+title, http.StatusFound) // what is statusfound
}

// this function will execute the code.
func executeCode(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/exec/"):]

	cmd := "/usr/local/go/bin/go"

	// execute the file and post the output
	out, err := exec.Command(cmd, "run", title+".go").Output()

	if err != nil {
		// TODO : if error then respond with error
		// message and give the link to go back to file.
		fmt.Print("could not execute")
		fmt.Fprint(w, err)
	} else {
		p := Page{Title: title, Output: out}

		htmlTemp, _ := template.ParseFiles("output.html")
		htmlTemp.Execute(w, p)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler) //handle all requests to the web root ("/") with handler
	http.HandleFunc("/edit/", editHandler) // handle the edit request.
	http.HandleFunc("/save/", saveHandler) // handle saving the page.
	http.HandleFunc("/exec/", executeCode) // handle the code execute.
	http.ListenAndServe(":9090", nil)      //function will block until program is terminated.
}

// Next : https://golang.org/doc/articles/wiki/# Saving Pages
