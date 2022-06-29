package main

import (
	"log"
	"net/http"

	"basicwebapp/controller"
	"basicwebapp/mapstore"
	"basicwebapp/router"
)

func main() {

	mux := router.Bookrouter()

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	controller.BL = *mapstore.NewBookLibrary()

	log.Print("Webserver started on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Print("Error occoured while starting webserver", err.Error())
	}
}
