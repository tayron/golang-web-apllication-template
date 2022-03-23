package controller

import (
	"net/http"
	"text/template"
	"time"
)

func ShowHomePage(res http.ResponseWriter, req *http.Request) {

	currentTime := time.Now()
	dateNow := currentTime.Format("02/01/2006 15:04:05")

	parameter := struct {
		Date string
	}{
		Date: dateNow,
	}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, parameter)
}
