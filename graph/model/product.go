package model

//Product Model
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// User        *User  `json:"user"`
	User string `json:"userId"`
}
