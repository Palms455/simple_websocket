package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(),
)

// Home views для домашней страницы
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

