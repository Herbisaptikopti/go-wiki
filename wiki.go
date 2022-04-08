package main

import (
	"fmt"
	"os"
)

type page struct {
	Title string
	Body  []byte
}

func (p *page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadpage(title string) *page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &page{Title: title, Body: body}
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func main() {
	p1 := &page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
