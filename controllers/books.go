package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"test-webapp/models"
	"encoding/hex"
)

const booksCollectionName string = "books" 

func BooksObjectIdsToString(books []models.Book) []models.Book {
	var newBooks []models.Book

	for _, book := range books {
		book.Id = hex.EncodeToString([]byte(book.Id))
		newBooks = append(newBooks, book)
	}
	return newBooks
}

func (mongoConnector MongoConnector) InsertBook(book *models.Book) {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Insert(book)

	if err != nil {
		panic(err)
	}
}

func (mongoConnector MongoConnector) FindBook(query bson.M) models.Book {
	session := mongoConnector.getSession()
	defer session.Close()

	fetchedBook := models.Book{}

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).One(&fetchedBook)

	if err != nil {
		panic(err)
	}

	return fetchedBook
}

func (mongoConnector MongoConnector) FindBooks(query bson.M) []models.Book {
	session := mongoConnector.getSession()
	defer session.Close()

	var fetchedBooks []models.Book

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).All(&fetchedBooks)

	if err != nil {
		panic(err)
	}

	return BooksObjectIdsToString(fetchedBooks)
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