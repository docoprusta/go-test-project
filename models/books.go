package models

type Book struct {  
    Id      string `json:"id" bson:"_id,omitempty"`
    ISBN    string   `json:"isbn"`
    Title   string   `json:"title"`
    Authors []string `json:"authors"`
    Price   float32   `json:"price"`
}
