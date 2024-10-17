package models

type Author struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name"`
}
