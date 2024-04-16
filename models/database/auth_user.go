package database

type User struct {
	Id        uint            `json:"id"`
	Firstname  string          `json:"firstname,omitempty"`
	Lastname  string          `json:"lastname,omitempty"`
	Email     string          `json:"email,omitempty"`
	Password  string          `json:"password,omitempty"`
}