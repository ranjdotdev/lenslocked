package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
	} 
	func contactHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:ranjdotdev@gmail.com\">ranjdotdev@gmail.com</a></p>")
}

func main() {
	r := chi.NewRouter();
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
    fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
