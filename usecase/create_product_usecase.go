package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product) error
	Retrive() ([]model.Product, error)
}

type createProductUseCase struct {
	repo repo.ProductRepo
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product) error {
	return c.repo.Add(newProduct)
}

func (c *createProductUseCase) Retrive() ([]model.Product, error) {
	return c.repo.Retrive()
}

func NewCreateProductUseCase(repo repo.ProductRepo) CreateProductUseCase {
	return &createProductUseCase{
		repo: repo,
	}
}
