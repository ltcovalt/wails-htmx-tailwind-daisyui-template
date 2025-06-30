package handlers

import (
	"html/template"
	"net/http"
)

type hx struct{}

var Hx hx

func (hx) Header(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/Header.html")
	if err != nil {
		return
	}

	data := struct{ Input string }{"Test input"}
	t.Execute(w, data)
}

func (hx) ThemeController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/ThemeController.html")
	if err != nil {
		return
	}

	t.Execute(w, nil)
}
