package models

type Book struct {  
    ISBN    string   `json:"isbn"`
    Title   string   `json:"title"`
    Authors []string `json:"authors"`
    Price   float32   `json:"price"`
}

const booksCollectionName string = "books" 

func (mongoConnector MongoConnector) InsertBook(book *Book) {
	session, err := mongoConnector.getSession()
	defer session.Close()

	if err != nil {
		panic(err)
	}

	c := session.DB(MongoConnectorVar.database).C(booksCollectionName)
	err = c.Insert(book)

}

func (mongoConnector MongoConnector) FindBook(book *Book) {
	session, err := mongoConnector.getSession()
	defer session.Close()

	if err != nil {
		panic(err)
	}

	c := session.DB(MongoConnectorVar.database).C(booksCollectionName)
	err = c.Insert(book)

}
