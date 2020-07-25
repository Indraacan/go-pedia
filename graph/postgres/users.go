package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//UserRepo for users
type UserRepo struct {
	DB *pg.DB
}

//GetUserByID controleer id
func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
