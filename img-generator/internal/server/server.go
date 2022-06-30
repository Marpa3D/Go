package server

import "net/http"

func Run() {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	http.ListenAndServe(":8080", nil)
}
