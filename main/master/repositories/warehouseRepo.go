package repositories

import "logisticApi/main/master/models"

type WarehouseRepository interface {
	GetAllWarehouse() ([]*models.Warehouse, error)
	GetWarehouseById(id string) (*models.Warehouse, error)
	Add([]*models.Warehouse) ([]*models.Warehouse, error)
	Update(idWarehouse string, Warehouse *models.Warehouse) (*models.Warehouse, error)
	Delete(idWarehouse string) error
}
