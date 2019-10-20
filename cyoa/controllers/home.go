package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("resources/templates/home.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	redirection := struct {
		Ref  string
		Text string
	}{
		Ref:  "/read/intro",
		Text: "start reading",
	}
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, redirection)
}
