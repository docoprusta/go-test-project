package main

import (
	"test-webapp/models"
)

func main() {
	book := models.Book {
		ISBN: "0131103628",
		Title: "C Programming Language, 2nd Edition",
		Authors: []string{"Brian W. Kernighan ", " Dennis M. Ritchie"},
		Price: 55.02,
	}

	models.MongoConnectorVar.InsertBook(&book)
}
