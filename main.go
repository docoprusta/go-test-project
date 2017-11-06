package main

import (
	"log"
	"test-webapp/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books", routes.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", routes.GetBook).Methods("GET")
	router.HandleFunc("/books", routes.PostBook).Methods("POST")
	router.HandleFunc("/book/{id}", routes.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", routes.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
