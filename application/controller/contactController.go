package controller

import (
	"net/http"
	"text/template"
)

func ShowHomePage(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, false)
}
