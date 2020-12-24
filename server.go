package main

import (
	"html/template"
	"net/http"
)

var tem *template.Template

func main() {
	http.HandleFunc("/", formpage)
	http.HandleFunc("/updated", updated)

	http.ListenAndServe(":8080", nil)

}

func formpage(w http.ResponseWriter, r *http.Request) {

	tem, _ := template.ParseFiles("form.html")

	tem.Execute(w, nil)

}

func updated(w http.ResponseWriter, r *http.Request) {

	tem, _ := template.ParseFiles("update.html")

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")
	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tem.Execute(w, d)

}
