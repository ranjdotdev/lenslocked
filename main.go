package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filename string){
	w.Header().Set("content-type", "text/html")
	t, err := template.ParseFiles("templates/"+ filename +".gohtml")
	if err != nil { 
		log.Println("Template parsing failed.\n", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return;
	}
	err = t.Execute(w, nil)
	if err != nil { 
		log.Println("Template execution failed.\n", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return;
	}
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
