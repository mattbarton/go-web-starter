package main

import (
	"net/http"
	"text/template"
	"log"
)

var config struct {
	liveTemplates bool
}

func init() {
	log.Print("init()")
	config.liveTemplates = true	
}

func main() {
	log.Print("main()")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", homeHandler)
	log.Print("starting")
	http.ListenAndServe(":8080", nil)
	log.Print("started")
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "home")
}

func compileTemplates() *template.Template {
	return template.Must(template.ParseGlob("./templates/*.html"))
}
// Precompile the templates
var precompiledTemplates = compileTemplates()

func renderTemplate(w http.ResponseWriter, tmpl string) {
	var templates *template.Template
	if config.liveTemplates {
		templates = compileTemplates()
	} else {
		templates = precompiledTemplates
	}
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
