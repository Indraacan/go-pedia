package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sony-nurdianto/go-pedia/graph/generated"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	_, err := r.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email is alredy exist")
	}

	_, err = r.UserRepo.GetUserByName(input.UserName)
	if err == nil {
		return nil, errors.New("user name is already used")
	}

	user := &model.User{
		UserName:  input.UserName,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hasing password : %v", err)
		return nil, errors.New("something went wrong")
	}

	//create verification code

	tx, err := r.UserRepo.DB.Begin()
	if err != nil {
		log.Printf("erro when creating transaction : %v", err)
		return nil, errors.New("something while wrong ")
	}

	defer tx.Rollback()

	if _, err := r.UserRepo.CreateUser(tx, user); err != nil {
		log.Printf("erro when creating user : %v", err)
		return nil, errors.New("something wrong")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("erro when commiting %v", err)
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Printf("erroe when generating token : %v", err)
		return nil, errors.New("something when wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name is to short")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description is to short")
	}

	product := &model.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	return r.ProductRepo.CreateProduct(product)
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
	product, err := r.ProductRepo.GetByID(id)
	if err != nil || product == nil {
		return nil, errors.New("product not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name to short")
		}
		product.Name = *input.Name

		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("Description to short")
		}
		product.Description = *input.Description
		didUpdate = true
	}

	if input.Price != nil {
		if *input.Price <= 0 {
			return nil, errors.New("price canot Zerro")
		}

		product.Price = *input.Price
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	product, err = r.ProductRepo.Update(product)
	if err != nil {
		return nil, fmt.Errorf("error while updateing %v", err)
	}

	return product, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	product, err := r.ProductRepo.GetByID(id)
	if err != nil || product == nil {
		return false, errors.New("product not exist")
	}

	err = r.ProductRepo.Delete(product)
	if err != nil {
		return false, fmt.Errorf("error while deleteing : %v", err)
	}

	return true, nil
}

func (r *productResolver) Users(ctx context.Context, obj *model.Product) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.User)
}

func (r *queryResolver) Products(ctx context.Context, filter *model.FilterProduct, limit *int, offset *int) ([]*model.Product, error) {
	return r.ProductRepo.GetProduct(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.GetUserByID(id)
}

func (r *userResolver) ProductID(ctx context.Context, obj *model.User) ([]*model.Product, error) {
	return r.ProductRepo.GetUserProduct(obj)
}

func (r *userResolver) UpdataeAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
