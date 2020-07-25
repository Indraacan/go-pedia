package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/go-pedia/graph/generated"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

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
