package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func Parse(filename string) (Template, error){
	tpl, err := template.ParseFiles("templates/"+ filename +".gohtml")
	if err != nil { 
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil;
}

func (t Template) Execute(w http.ResponseWriter, data interface{}){
	w.Header().Set("content-type", "text/html")
	err := t.HTMLTpl.Execute(w, data)
	if err != nil { 
		log.Println("Template execution failed.\n", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return;
	}
}