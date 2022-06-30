package router

import (
	"basicwebapp/controller"

	"net/http"

	"github.com/gorilla/mux"
)

func Bookrouter(api *controller.BookAPIController) *mux.Router {
	mux := mux.NewRouter()

	mux.Handle("/", http.HandlerFunc(api.Welcome)).Methods("GET")
	mux.Handle("/book", http.HandlerFunc(api.GetAll)).Methods("GET")
	mux.Handle("/book/{id}", http.HandlerFunc(api.Get)).Methods("GET")
	mux.Handle("/book", http.HandlerFunc(api.Post)).Methods("POST")
	mux.Handle("/book/{id}", http.HandlerFunc(api.Put)).Methods("PUT")
	mux.Handle("/book/{id}", http.HandlerFunc(api.Delete)).Methods("DELETE")
	return mux
}
