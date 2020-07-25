package model

//User Model
type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	// Products []*Product `json:"products"`
}
