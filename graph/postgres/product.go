package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//ProductRepo Tank for object
type ProductRepo struct {
	DB *pg.DB
}

//GetProduct all
func (p *ProductRepo) GetProduct() ([]*model.Product, error) {

	var products []*model.Product
	err := p.DB.Model(&products).Select()
	if err != nil {
		return nil, err
	}

	return products, nil
}

//CreateProduct Post i think
func (p *ProductRepo) CreateProduct(product *model.Product) (*model.Product, error) {
	_, err := p.DB.Model(product).Returning("*").Insert()

	return product, err
}

//GetByID Product
func (p *ProductRepo) GetByID(id string) (*model.Product, error) {
	var product model.Product

	err := p.DB.Model(&product).Where("id = ?", id).First()

	return &product, err

}

//Update product
func (p *ProductRepo) Update(product *model.Product) (*model.Product, error) {
	_, err := p.DB.Model(product).Where("id = ? ", product.ID).Update()
	return product, err

}

//Delete product
func (p *ProductRepo) Delete(product *model.Product) error {
	_, err := p.DB.Model(product).Where("id = ?", product.ID).Delete()
	return err
}

//GetUserProduct

func (p *ProductRepo) GetUserProduct(user *model.User) ([]*model.product, error) {
	var products []*model.Product
	err := p.DB.Model(&products).Where("user = ?", user.ID).Select()
	return products, err
}
