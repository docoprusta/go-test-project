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

func BookObjectIdToString(book *models.Book) {
	book.Id = hex.EncodeToString([]byte(book.Id))
}

func (mongoConnector MongoConnector) InsertBook(book *models.Book) error {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Insert(book)

	return err
}

func (mongoConnector MongoConnector) FindBook(query bson.M) (models.Book, error) {
	session := mongoConnector.getSession()
	defer session.Close()

	var fetchedBook models.Book

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).One(&fetchedBook)

	BookObjectIdToString(&fetchedBook)

	return fetchedBook, err
}

func (mongoConnector MongoConnector) FindBooks(query bson.M) ([]models.Book, error) {
	session := mongoConnector.getSession()
	defer session.Close()

	var fetchedBooks []models.Book

	c := session.DB(mongoConnector.database).C(booksCollectionName)
	err := c.Find(query).All(&fetchedBooks)

	return BooksObjectIdsToString(fetchedBooks), err
}

func (mongoConnector MongoConnector) UpdateBook(query bson.M, change bson.M) error {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)

	err := c.Update(query, change)
	return err
}

func (mongoConnector MongoConnector) RemoveBook(query bson.M) error {
	session := mongoConnector.getSession()
	defer session.Close()

	c := session.DB(mongoConnector.database).C(booksCollectionName)

	err := c.Remove(query)
	return err
}
