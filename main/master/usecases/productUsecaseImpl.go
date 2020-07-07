package usecases

import (
	"logisticApi/main/master/models"
	"logisticApi/main/master/repositories"
	"logisticApi/utils"
)

type ProductUsecaseImpl struct {
	productRepo repositories.ProductRepository
}

func (p *ProductUsecaseImpl) GetAll() ([]*models.Product, error) {
	products, err := p.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (p *ProductUsecaseImpl) GetById(id string) (*models.Product, error) {
	product, err := p.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *ProductUsecaseImpl) DeleteProduct(id string) error {
	err := s.productRepo.Delete(id)
	return err
}
func (s *ProductUsecaseImpl) AddProduct(data []*models.Product) ([]*models.Product, error) {
	err := utils.ValidateInput(data)
	product, err := s.productRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *ProductUsecaseImpl) UpdateProduct(id string, data *models.Product) (*models.Product, error) {
	err := utils.ValidateInput(data)
	product, err := s.productRepo.Update(id, data)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func InitProductUsecaseImpl(ProductRepo repositories.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{ProductRepo}
}
