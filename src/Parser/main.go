// Простой парсер всех ссылок по указанному URL
package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Посещаю ресурс:", r.URL)
	})

	c.Visit("https://forbes.com/")
}
