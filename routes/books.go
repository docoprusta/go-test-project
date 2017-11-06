package routes

import (
	"test-webapp/models"
	"io/ioutil"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"test-webapp/controllers"
	"encoding/json"
    "net/http"
)

var mongoConnector controllers.MongoConnector

func internalServerError(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(writer).Encode(bson.M{"msg": "Internal server error"})
}

func notFound(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusNotFound)
    json.NewEncoder(writer).Encode(bson.M{"msg": "Not found"})
}

func badRequest(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(writer).Encode(bson.M{"msg": "Bad request"})
}

func ok(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(bson.M{"msg": "Ok"})
}

func GetBooks(writer http.ResponseWriter, request *http.Request) {
    mongoConnector.InitMongoValues()
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    books, err := mongoConnector.FindBooks(bson.M{})
    if err != nil {
        internalServerError(writer, request)
    }
    json.NewEncoder(writer).Encode(books)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
    
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    params := mux.Vars(request)
    if len(params["id"]) != 24 {
        badRequest(writer, request)
        return
    }

    mongoConnector.InitMongoValues()

    book, err := mongoConnector.FindBook(bson.M{"_id": bson.ObjectIdHex(params["id"])})

    if err != nil {
        if (err.Error() == "not found") {
            notFound(writer, request)
            return
        } else {
            writer.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(writer).Encode(bson.M{"msg": "Internal server error"})
            return
        }
    }
    
    json.NewEncoder(writer).Encode(book)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
    
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    params := mux.Vars(request)
    if len(params["id"]) != 24 {
        badRequest(writer, request)
        return
    }

    mongoConnector.InitMongoValues()

    err := mongoConnector.RemoveBook(bson.M{"_id": bson.ObjectIdHex(params["id"])})

    if err != nil {
        if (err.Error() == "not found") {
            notFound(writer, request)
            return
        } else {
            writer.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(writer).Encode(bson.M{"msg": "Internal server error"})
            return
        }
    }

    ok(writer, request)
}

func PostBook(writer http.ResponseWriter, request *http.Request) {
    mongoConnector.InitMongoValues()
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    body, err := ioutil.ReadAll(request.Body)
    
	if err != nil {
        internalServerError(writer, request)
		return
    }

    var book models.Book

    err = json.Unmarshal([]byte(body), &book)

    if err != nil {
        internalServerError(writer, request)
        return
    }

    mongoConnector.InsertBook(&book)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
    
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    params := mux.Vars(request)
    if len(params["id"]) != 24 {
        badRequest(writer, request)
        return
    }

    mongoConnector.InitMongoValues()
    body, err := ioutil.ReadAll(request.Body)

    m := bson.M{}
    if err := json.Unmarshal(body, &m); err != nil {
        panic(err)
    }

    err = mongoConnector.UpdateBook(bson.M{"_id": bson.ObjectIdHex(params["id"])}, m)

    if err != nil {
        if (err.Error() == "not found") {
            notFound(writer, request)
            return
        } else {
            writer.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(writer).Encode(bson.M{"msg": "Internal server error"})
            return
        }
    }
    
    ok(writer, request)
}