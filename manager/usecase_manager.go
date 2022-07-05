package manager

import (
	"go-api-with-gin2/usecase"
)

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUseCase
	ListProductUseCase() usecase.ListProductUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ListProductUseCase() usecase.ListProductUseCase {
	return usecase.NewListProductUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
