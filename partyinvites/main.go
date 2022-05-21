package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "sanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layuot.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}
func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}
