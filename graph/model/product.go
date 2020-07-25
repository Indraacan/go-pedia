package model

//Product Model
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	// User        *User  `json:"userid"`
	User string `json:"users"`
}
