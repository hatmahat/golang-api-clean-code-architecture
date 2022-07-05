package manager

import (
	"go-api-with-gin2/repo"
)

type RepoManager interface {
	// kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repo.ProductRepo
}

type repoManager struct {
	// productRepo repo.ProductRepo
	infra Infra
}

func (r *repoManager) ProductRepo() repo.ProductRepo {
	return repo.NewProductRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	// productRepo := repo.NewProductRepo()
	return &repoManager{
		infra: infra,
	}
}
