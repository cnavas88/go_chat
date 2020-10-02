package main

import (
	"fmt"
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

var port string = ":8080"
var templatesFolder string = "templates"

func main() {
	t := templateHandler{filename: "chat.html"}

	http.HandleFunc("/", t.ServeHTTP)

	// Start the web server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If once is called multiple times, only the
	// first call invoke func()
	t.once.Do(func() {
		fmt.Println(filepath.Join(templatesFolder, t.filename))
		// Must method wraps a call to returning template, err parameters
		// if there are an error, must execute panic function
		t.templ = template.Must(template.ParseFiles(filepath.Join(templatesFolder, t.filename)))
	})

	t.templ.Execute(w, nil)
}
