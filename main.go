package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ranjdotdev/lenslocked/controllers"
	"github.com/ranjdotdev/lenslocked/views"
)

func main() {
	r := chi.NewRouter();

	tpl, err := views.Parse("home"); if err != nil {panic(err)}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("contact"); if err != nil {panic(err)}
	r.Get("/contact", controllers.StaticHandler(tpl))
	
	tpl, err = views.Parse("faq"); if err != nil {panic(err)}
	r.Get("/faq", controllers.StaticHandler(tpl))

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
