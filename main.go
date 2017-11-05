package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"test-webapp/models"
)

func main() {
	book1 := models.Book {
		ISBN: "0131103628",
		Title: "C Programming Language, 2nd Edition",
		Authors: []string{"Brian W. Kernighan ", " Dennis M. Ritchie"},
		Price: 55.02,
	}

	book2 := models.Book {
		ISBN: "1593072937",
		Title: "The Hard Goodbye (Sin City)",
		Authors: []string{"Frank Miller"},
		Price: 12.86,
	}

	var mongoConnector models.MongoConnector
	mongoConnector.InitMongoValues()

	mongoConnector.InsertBook(&book1)
	mongoConnector.InsertBook(&book2)

	fetchedBook := mongoConnector.FindBook(bson.M{"price": bson.M{"$lte": 30}})
	// fetchedBooks := mongoConnector.FindBooks(bson.M{"price": bson.M{"$lte": 30}})

	mongoConnector.UpdateBook(bson.M{"isbn": "1593072937"}, bson.M{"isbn": "1"})

	mongoConnector.RemoveBook(bson.M{"isbn": "1"})
	// fetchedBook := mongoConnector.FindBook(bson.M{"isbn": "1"})
	fmt.Println(fetchedBook.Title)
}
