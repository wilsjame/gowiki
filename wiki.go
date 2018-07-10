package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page defined by a title and body.
type Page struct {
	Title string
	Body []byte	// Type expected by io libraries.
}

// Persistent storage.
// This is method name 'save' that takes as its receiver 'p',
// a pointer to 'Page'. It takes no parameters, and returns a
// value of type 'error'.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Allow user to view a wiki page.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



