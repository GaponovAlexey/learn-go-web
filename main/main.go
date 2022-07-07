package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html")
	if err != nil {
		fmt.Println(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "create", nil)
}

func main() {
	HandleFunc()
}

func HandleFunc() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":3000", nil)
}
