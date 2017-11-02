package models

type Book struct {  
    ISBN    string   `json:"isbn"`
    Title   string   `json:"title"`
    Authors []string `json:"authors"`
    Price   int   `json:"price"`
}

func (mongoConnector MongoConnector) insertBook(book Book) {
	session, err := mongoConnector.getSession()

	if err != nil {
		panic(err)
	}
	c := session.DB("test").C("people")
	err = c.Insert(&Book{ISBN: "12345", Title: "asd", Authors: []string{"",""}, Price: 45})
}
