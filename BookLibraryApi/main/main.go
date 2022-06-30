package main

import (
	"log"
	"net/http"

	"basicwebapp/controller"
	"basicwebapp/mapstore"
	"basicwebapp/router"
)

func main() {

	apiController := &controller.BookAPIController{
		BL: mapstore.NewBookLibrary(),
	}

	bookMux := router.Bookrouter(apiController)

	server := &http.Server{
		Addr:    ":8000",
		Handler: bookMux,
	}

	log.Print("Webserver started on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Print("Error occoured while starting webserver", err.Error())
	}
}
