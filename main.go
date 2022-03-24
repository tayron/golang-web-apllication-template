package main

import (
	"fmt"
	"googleauth/application"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

var clientId string = "845986572673-lcte5kb6hipc4dcg10fc79j4s9emca61.apps.googleusercontent.com"
var clientSecred string = "niLaSNdNsG9Vq19YgA_Dmams"

func main() {

	// Configurando timezone
	var timeZone = time.FixedZone("America/Sao_Paulo", -3*3600)
	time.Local = timeZone

	key := "webapplication" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30    // 30 days
	isProd := false         // Set to true when serving over https

	if os.Getenv("ENVIROMENT") != "developer" {
		isProd = true
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	url := fmt.Sprintf("%s:%s", os.Getenv("URL"), os.Getenv("PORT"))

	log.Println("listening on " + url)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), application.GetRoutes()))
}
