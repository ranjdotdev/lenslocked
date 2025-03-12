package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ranjdotdev/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filename string){
	tpl, err := views.Parse(filename)
	if err != nil { 
		log.Println("Template parsing failed.\n", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return;
	}
	tpl.Execute(w, nil)
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home")
}
func contactHandler (w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact")
}
func faqHandler (w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "faq")
}

func main() {
	r := chi.NewRouter();
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
    fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
