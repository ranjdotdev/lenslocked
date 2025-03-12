package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User {
		Name: "ranj",
		Bio: `<script>alert("Haha, you have been hx0r3d");</script>`,
	} 
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
	fmt.Println()
} 