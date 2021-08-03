package application

import (
	"googleauth/application/controller"
	"googleauth/application/util"
	"net/http"

	"github.com/gorilla/pat"
)

func GetRoutes() *pat.Router {
	router := pat.New()

	diretorioAplicacao := util.GetApplicationRootPath()

	publicJs := diretorioAplicacao + "/public/js/"
	js := http.StripPrefix("/js/", http.FileServer(http.Dir(publicJs)))
	router.PathPrefix("/js").Handler(js)

	publicCss := diretorioAplicacao + "/public/css/"
	css := http.StripPrefix("/css/", http.FileServer(http.Dir(publicCss)))
	router.PathPrefix("/css").Handler(css)

	publicAssets := diretorioAplicacao + "/public/assets/"
	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir(publicAssets)))
	router.PathPrefix("/assets").Handler(assets)

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		controller.ShowHomePage(res, req)
	})

	return router
}
