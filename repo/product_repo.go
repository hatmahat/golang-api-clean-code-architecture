package repo

import (
	"go-api-with-gin2/model"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Add(newProduct *model.Product) error
	Retrive() ([]model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func (p *productRepo) Retrive() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepo) Add(newProduct *model.Product) error {
	err := p.db.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	repo := new(productRepo)
	repo.db = db
	return repo
}
