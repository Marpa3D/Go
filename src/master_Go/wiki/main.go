// Web app. Wiki
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Структура страниц wiki. Описание, как эти данные будут храниться в памяти
type Page struct {
	Title string
	Body  []byte
}

// Save() - метод c приемником-указателем на структуру Page записывает (сохраняет) данные в файл типа *.txt
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // 0600 - код, указывает,  толко для
	// что файл для записи и чтения, только для текущего пользователя
}

// Загрузка страниц
func LoadFile(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Обработчик запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, я люблю %s!", r.URL.Path[1:])
}

func main() {
	p1 := &Page{
		Title: "testPage",
		Body:  []byte("Это пример тестовой страницы Wiki"),
	}
	p1.Save()

	p2, _ := LoadFile("testPage")
	fmt.Println(string(p2.Body))

	// Обраотчик и сервер
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
