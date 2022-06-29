package router

import (
	"basicwebapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Bookrouter() *mux.Router {
	mux := mux.NewRouter()
	mux.Handle("/", http.HandlerFunc(controller.Welcome)).Methods("GET")
	mux.Handle("/book", http.HandlerFunc(controller.GetAll)).Methods("GET")
	mux.Handle("/book/{id}", http.HandlerFunc(controller.Get)).Methods("GET")
	mux.Handle("/book", http.HandlerFunc(controller.Post)).Methods("POST")
	mux.Handle("/book/{id}", http.HandlerFunc(controller.Put)).Methods("PUT")
	mux.Handle("/book/{id}", http.HandlerFunc(controller.Delete)).Methods("DELETE")
	return mux
}
