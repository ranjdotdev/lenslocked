package main

import (
	"fmt"
	"net/http"
)

func handleFunc (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
} 

func main() {

	http.HandleFunc("/", handleFunc)
    fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
