package router

import (
	"url-shortner/controller"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/url", controller.CreateURL).Methods("POST")
	r.HandleFunc("/urls", controller.GetAllURLs).Methods("GET")
	r.HandleFunc("/url/{id}", controller.GetURLByID).Methods("GET")

	return r
}
