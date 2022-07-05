package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListProductUseCase interface {
	Retrive() ([]model.Product, error)
}

type listProductUseCase struct {
	repo repo.ProductRepo
}

func (c *listProductUseCase) Retrive() ([]model.Product, error) {
	return c.repo.Retrive()
}

func NewListProductUseCase(repo repo.ProductRepo) ListProductUseCase {
	return &listProductUseCase{
		repo: repo,
	}
}
