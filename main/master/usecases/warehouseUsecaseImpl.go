package usecases

import (
	"logisticApi/main/master/models"
	"logisticApi/main/master/repositories"
	"logisticApi/utils"
)

type WarehouseUsecaseImpl struct {
	warehouseRepo repositories.WarehouseRepository
}

func (w *WarehouseUsecaseImpl) GetAll() ([]*models.Warehouse, error) {
	warehouses, err := w.warehouseRepo.GetAllWarehouse()
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}
func (w *WarehouseUsecaseImpl) GetById(id string) (*models.Warehouse, error) {
	warehouse, err := w.warehouseRepo.GetWarehouseById(id)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}
func (w *WarehouseUsecaseImpl) DeleteWarehouse(id string) error {
	err := w.warehouseRepo.Delete(id)
	return err
}
func (w *WarehouseUsecaseImpl) AddWarehouse(data []*models.Warehouse) ([]*models.Warehouse, error) {
	err := utils.ValidateInput(data)
	warehouse, err := w.warehouseRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}
func (w *WarehouseUsecaseImpl) UpdateWarehouse(id string, data *models.Warehouse) (*models.Warehouse, error) {
	err := utils.ValidateInput(data)
	warehouse, err := w.warehouseRepo.Update(id, data)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}
func InitWarehouseUsecaseImpl(warehouseRepo repositories.WarehouseRepository) WarehouseUsecase {
	return &WarehouseUsecaseImpl{warehouseRepo}
}
