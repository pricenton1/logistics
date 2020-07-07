package models

type Warehouse struct {
	IdWarehouse       string `json:"id"`
	WarehouseName     string `json:"warehouseName"`
	WarehouseLocation string `json:"warehouseLocation"`
	WarehouseCapacity string `json:"warehouseCapacity"`
	WarehouseType     WarehouseType
}
