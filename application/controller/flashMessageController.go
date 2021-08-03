package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

type FlashMessage struct {
	Message string
	Type    string
}

type FlashMessageList struct {
	Messages []FlashMessage
}

func SetFlashMessage(res http.ResponseWriter, req *http.Request, message string, messageType string) {
	session, _ := gothic.Store.Get(req, "flash-session")

	session.AddFlash(message, messageType)

	err := session.Save(req, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func GetFlashMessage(res http.ResponseWriter, req *http.Request) FlashMessageList {
	listFlashMessage := FlashMessageList{}

	session, _ := gothic.Store.Get(req, "flash-session")

	sessionFlashMessageSuccess := session.Flashes("success")
	if sessionFlashMessageSuccess != nil {
		flashMessage := FlashMessage{
			Message: fmt.Sprintf("%v", sessionFlashMessageSuccess[0]),
			Type:    "success",
		}

		listFlashMessage.Messages = append(listFlashMessage.Messages, flashMessage)
	}

	sessionFlashMessageError := session.Flashes("error")
	if sessionFlashMessageError != nil {
		flashMessage := FlashMessage{
			Message: fmt.Sprintf("%v", sessionFlashMessageError[0]),
			Type:    "error",
		}

		listFlashMessage.Messages = append(listFlashMessage.Messages, flashMessage)
	}

	clearFlashMessage(res, req)
	return listFlashMessage

}

func clearFlashMessage(res http.ResponseWriter, req *http.Request) {
	session, _ := gothic.Store.Get(req, "flash-session")
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	err := session.Save(req, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
