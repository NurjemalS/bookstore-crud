package models

type Book struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
