package usecases

import "logisticApi/main/master/models"

type WarehouseUsecase interface {
	GetAll() ([]*models.Warehouse, error)
	GetById(id string) (*models.Warehouse, error)
	AddWarehouse([]*models.Warehouse) ([]*models.Warehouse, error)
	UpdateWarehouse(idWarehouse string, Warehouse *models.Warehouse) (*models.Warehouse, error)
	DeleteWarehouse(idWarehouse string) error
}
