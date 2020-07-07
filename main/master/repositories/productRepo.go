package repositories

import "logisticApi/main/master/models"

type ProductRepository interface {
	GetAllProduct() ([]*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	Add([]*models.Product) ([]*models.Product, error)
	Update(idProduct string, product *models.Product) (*models.Product, error)
	Delete(idProduct string) error
}
