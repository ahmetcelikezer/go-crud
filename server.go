package main

import (
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

type _page struct {
	Title       string
	Description string
}

func index(write http.ResponseWriter, request *http.Request) {
	data := _page{"Go Web Server", "Website build with GOLang"}

	templatePath := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(templatePath)

	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(write, data); err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
}
