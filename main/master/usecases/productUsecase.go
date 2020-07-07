package usecases

import "logisticApi/main/master/models"

type ProductUsecase interface {
	GetAll() ([]*models.Product, error)
	GetById(id string) (*models.Product, error)
	AddProduct([]*models.Product) ([]*models.Product, error)
	UpdateProduct(idProduct string, product *models.Product) (*models.Product, error)
	DeleteProduct(idProduct string) error
}
