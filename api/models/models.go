package models

type Todos struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Status bool   `json:"status"`
}
