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
	hTmp, err := template.ParseFiles("templates/"+ filename)
	if err != nil { 
		log.Println("Template parsing failed.\n", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return;
	}
	err = hTmp.Execute(w, nil)
	if err != nil { 
		log.Println("Template execution failed.\n", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return;
	}
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home.gohtml")
}
func contactHandler (w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact.gohtml")
}

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	userID := chi.URLParam(r, "userID")
// 	w.Write([]byte(fmt.Sprintf("hi %v", userID)))
//   }

func main() {
	r := chi.NewRouter();
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	// r.Get("/contact/{userID}", userHandler)
    fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
