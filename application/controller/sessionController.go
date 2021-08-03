package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func setSession(res http.ResponseWriter, req *http.Request, name string, value string) {
	session, _ := gothic.Store.Get(req, "webapplication")
	session.Values[name] = value

	err := session.Save(req, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func getSession(req *http.Request, name string) string {
	session, _ := gothic.Store.Get(req, "webapplication")
	dado, err := session.Values[name]

	if err == false {
		return ""
	}

	return dado.(string)
}

func clearSession(res http.ResponseWriter, req *http.Request) {
	session, _ := gothic.Store.Get(req, "webapplication")
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	err := session.Save(req, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
