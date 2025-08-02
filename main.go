package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Variable created to reference the html index page

func main() {
	http.HandleFunc("/", index) // (router , function of handle)
	http.ListenAndServe(":33", nil)
	// *:8000* Port used to host the HTTP server
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil) // (writer referenced by *w* , indicator of the page that should be rendered)
	// Call the template variable and execute the template page
}
