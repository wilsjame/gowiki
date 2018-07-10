package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}









