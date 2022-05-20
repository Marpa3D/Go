package main

import (
	"fmt"
	"html/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAtteend bool
}
var response = make([]*Rsvp, 0, 10)
var templates =make(map[string]*template.Template, 3)

func loadTemlates()  {
	templateNames = [5]string{"welcome", "form", "sanks", "sorry", "list"}
	for index, name = range templates{
		t, err = template.ParseFiles("layuot.html", name + ".html")
		if err == nil{
			templates[name] = t
			fmt.Println("Загружен шаблон", index, name)
		} else {
			panic(err)
		}
	}
}

func main()  {
	fmt.Println("Поехали!")
	loadTemlates()
}
