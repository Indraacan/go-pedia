package model

//User Model
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	// ProductID []*Product `json:"product_id"`
}
