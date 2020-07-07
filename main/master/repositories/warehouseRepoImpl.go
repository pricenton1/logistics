package repositories

import (
	"database/sql"
	"logisticApi/main/master/models"
)

type WarehouseRepoImpl struct {
	db *sql.DB
}

func (w *WarehouseRepoImpl) GetAllWarehouse() ([]*models.Warehouse, error) {
	var warehouses []*models.Warehouse
	query := "SELECT * FROM warehouse w JOIN warehouse_type wt ON w.id_type=wt.id"
	rows, err := w.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		warehouse := models.Warehouse{}
		err := rows.Scan(&warehouse.IdWarehouse, &warehouse.WarehouseName, &warehouse.WarehouseLocation, &warehouse.WarehouseCapacity, &warehouse.WarehouseType.IdType, &warehouse.WarehouseType.IdType, &warehouse.WarehouseType.TypeWarehouse)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, &warehouse)
	}
	return warehouses, nil
}
func (w *WarehouseRepoImpl) GetWarehouseById(id string) (*models.Warehouse, error) {
	row := w.db.QueryRow("SELECT * FROM warehouse w JOIN warehouse_type wt ON w.id_type=wt.id WHERE w.id_warehouse=?", id)
	var warehouse = models.Warehouse{}
	err := row.Scan(&warehouse.IdWarehouse, &warehouse.WarehouseName, &warehouse.WarehouseLocation, &warehouse.WarehouseCapacity, &warehouse.WarehouseType.IdType, &warehouse.WarehouseType.IdType, &warehouse.WarehouseType.TypeWarehouse)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil

}
func (w *WarehouseRepoImpl) Add(warehouses []*models.Warehouse) ([]*models.Warehouse, error) {
	for _, warehouse := range warehouses {

		tx, err := w.db.Begin()
		if err != nil {
			return nil, err
		}
		res, err := tx.Exec("INSERT INTO warehouse VALUES (?,?,?,?,?)", &warehouse.IdWarehouse, &warehouse.WarehouseName, &warehouse.WarehouseLocation, &warehouse.WarehouseCapacity, &warehouse.WarehouseType.IdType)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		_, err = tx.Exec("INSERT INTO warehouse_type VALUES (?,?)", id, &warehouse.WarehouseType.TypeWarehouse)

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		return nil, tx.Commit()
	}

	return warehouses, nil
}
func (w *WarehouseRepoImpl) Delete(id string) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM warehouse WHERE id_warehouse=?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}
func (w *WarehouseRepoImpl) Update(id string, warehouse *models.Warehouse) (*models.Warehouse, error) {
	tx, err := w.db.Begin()
	_, err = tx.Exec("UPDATE product SET warehouse_name=?,warehouse_location=?,warehouse_capacity=?,id_type=? WHERE id_warehouse=?", &warehouse.WarehouseName, &warehouse.WarehouseLocation, &warehouse.WarehouseCapacity, &warehouse.WarehouseType.IdType, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return warehouse, tx.Commit()
}
func InitWarehouseRepoImpl(db *sql.DB) WarehouseRepository {
	return &WarehouseRepoImpl{db}
}
