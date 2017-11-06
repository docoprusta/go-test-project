package routes

import (
	"github.com/gorilla/mux"
	"test-webapp/models"
	"gopkg.in/mgo.v2/bson"
	"test-webapp/controllers"
	"encoding/json"
    "net/http"
)

var mongoConnector controllers.MongoConnector

func GetBooks(writer http.ResponseWriter, request *http.Request) {
    mongoConnector.InitMongoValues()
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    books := mongoConnector.FindBooks(bson.M{})
    json.NewEncoder(writer).Encode(books)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
    var books []models.Book

    mongoConnector.InitMongoValues()
    params := mux.Vars(request)
    for _, item := range books {
        if item.Id == params["id"] {
            json.NewEncoder(writer).Encode(item)
            return
        }
    }
    json.NewEncoder(writer).Encode(&models.Book{})
}
