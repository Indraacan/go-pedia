package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/sony-nurdianto/go-pedia/graph/generated"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

var products = []*model.Product{
	{
		ID:          "1",
		Name:        "Product",
		Description: "product Description",
		Price:       20000,
		// User:        "1",
	},
	{
		ID:          "1",
		Name:        "Product",
		Description: "product Description",
		Price:       2000,
		// User:        "2",
	},
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name is to short")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description is to short")
	}

	product := &model.Product{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	return r.ProductRepo.CreateProduct(product)
}

func (r *productResolver) Users(ctx context.Context, obj *model.Product) (*model.User, error) {
	return r.UserRepo.GetUserByID(obj.ID)
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return r.ProductRepo.GetProduct()
}

func (r *userResolver) Products(ctx context.Context, obj *model.User) ([]*model.Product, error) {
	var m []*model.Product

	for _, products := range products {

		if products.User == obj.ID {
			m = append(m, products)
		}

	}

	return m, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
