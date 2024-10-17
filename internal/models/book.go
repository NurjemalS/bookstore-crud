package models

type Book struct {
	ID       int     `json:"id" form:"id"`
	Title    string  `json:"title"`
	Author   Author  `json:"author"`
	AuthorID int     `json:"author_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	FilePath string  `json:"file_path"` // file is optonal, different endpoint for uploading seperately
}
