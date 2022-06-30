package controller

import (
	"basicwebapp/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BookAPIController struct {
	BL model.BookController
}

// Handler Function for default route - Healthcheck
func (b *BookAPIController) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := model.HealthCheck{Data: "Welcome to Book Library Application"}
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func (b *BookAPIController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.HealthCheck{Data: err.Error()})
		return
	}

	err = b.BL.Add(book)
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)

}

func (b *BookAPIController) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := b.BL.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.HealthCheck{Data: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)

}

func (b *BookAPIController) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := b.BL.GetBook(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.HealthCheck{Data: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)

}

func (b *BookAPIController) Put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("Id:", id)
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	err := b.BL.Update(id, book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.HealthCheck{Data: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func (b *BookAPIController) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := b.BL.Delete(id)
	if err != nil {
		json.NewEncoder(w).Encode(model.HealthCheck{Data: err.Error()})
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(model.HealthCheck{Data: "Book Deleted"})
	w.WriteHeader(http.StatusOK)
}
