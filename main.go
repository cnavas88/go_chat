package main

import (
	"html/template" // to load templates when init program
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template // variable to store templates
}

func main() {
	t := templateHandler{filename: "chat.html"}

	http.HandleFunc("/", t.ServeHTTP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If once is called multiple times, only the
	// first call invoke func()
	t.once.Do(func() {
		// Must method wraps a call to returning template, err parameters
		// if there are an error, must execute panic function
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}
