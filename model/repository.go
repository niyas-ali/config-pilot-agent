package model

type Repository struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Branch string `json:"branch"`
}
