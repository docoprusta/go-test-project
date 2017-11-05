package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Book struct {  
    ISBN    string   `json:"isbn"`
    Title   string   `json:"title"`
    Authors []string `json:"authors"`
    Price   float32   `json:"price"`
}

const booksCollectionName string = "books" 

func (mongoConnector MongoConnector) InsertBook(book *Book) {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Insert(book)

	if err != nil {
		panic(err)
	}
}

func (mongoConnector MongoConnector) FindBook(query bson.M) Book {
	session := mongoConnector.getSession()
	defer session.Close()

	fetchedBook := Book{}

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).One(&fetchedBook)

	if err != nil {
		panic(err)
	}

	return fetchedBook
}

func (mongoConnector MongoConnector) FindBooks(query bson.M) []Book {
	session := mongoConnector.getSession()
	defer session.Close()

	var fetchedBooks []Book

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).All(&fetchedBooks)

	if err != nil {
		panic(err)
	}

	return fetchedBooks
}

func (mongoConnector MongoConnector) UpdateBook(query bson.M, change bson.M) {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)

	err := c.Update(query, change)
	if err != nil {
		panic(err)
	}
}

func (mongoConnector MongoConnector) RemoveBook(query bson.M) {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)

	err := c.Remove(query)
	if err != nil {
		panic(err)
	}
}