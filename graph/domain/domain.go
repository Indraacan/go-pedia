package domain

import "github.com/sony-nurdianto/go-pedia/graph/postgres"

//Domain Struct same function to store Data from DataBase
type Domain struct {
	UserRepo    postgres.UserRepo
	ProductRepo postgres.ProductRepo
}

func NewDomain(userRepo postgres.UserRepo, productRepo postgres.ProductRepo) *Domain {
	return &Domain{UserRepo: userRepo, ProductRepo: productRepo}
}


type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnerShip(o Ownable , user *model.User) bool {
	return o.IsOwner(user)
}