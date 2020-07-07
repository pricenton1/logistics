package repositories

import (
	"database/sql"
	"logisticApi/main/master/models"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (p *ProductRepoImpl) GetAllProduct() ([]*models.Product, error) {
	var products []*models.Product
	query := "SELECT * FROM product p JOIN warehouse w ON p.warehouse_id=w.id_warehouse JOIN warehouse_type wt ON w.id_type=wt.id"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.IdProduct, &product.ProductName, &product.Qty, &product.Type, &product.ProductWarehouse.IdWarehouse, &product.ProductWarehouse.IdWarehouse, &product.ProductWarehouse.WarehouseName, &product.ProductWarehouse.WarehouseLocation, &product.ProductWarehouse.WarehouseCapacity, &product.ProductWarehouse.WarehouseType.IdType, &product.ProductWarehouse.WarehouseType.IdType, &product.ProductWarehouse.WarehouseType.TypeWarehouse)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
func (p *ProductRepoImpl) GetProductById(id string) (*models.Product, error) {
	row := p.db.QueryRow("SELECT * FROM product p JOIN warehouse w ON p.warehouse_id=w.id_warehouse JOIN warehouse_type wt ON w.id_type=wt.id WHERE p.id_product=?", id)
	var product = models.Product{}
	err := row.Scan(&product.IdProduct, &product.ProductName, &product.Qty, &product.Type, &product.ProductWarehouse.IdWarehouse, &product.ProductWarehouse.IdWarehouse, &product.ProductWarehouse.WarehouseName, &product.ProductWarehouse.WarehouseLocation, &product.ProductWarehouse.WarehouseCapacity, &product.ProductWarehouse.WarehouseType.IdType, &product.ProductWarehouse.WarehouseType.IdType, &product.ProductWarehouse.WarehouseType.TypeWarehouse)
	if err != nil {
		return nil, err
	}
	return &product, nil

}
func (p *ProductRepoImpl) Add(products []*models.Product) ([]*models.Product, error) {
	for _, product := range products {

		tx, err := p.db.Begin()
		if err != nil {
			return nil, err
		}
		_, err = tx.Exec("INSERT INTO product VALUES (?,?,?,?)", &product.IdProduct, &product.ProductName, &product.Qty, &product.Type, &product.ProductWarehouse.IdWarehouse)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// id, err := res.LastInsertId()
		// if err != nil {
		// 	tx.Rollback()
		// 	return nil, err
		// }

		// _, err = tx.Exec("INSERT INTO warehouse VALUES ()")

		// if err != nil {
		// 	tx.Rollback()
		// 	return nil, err
		// }
		return nil, tx.Commit()
	}

	return products, nil
}
func (p *ProductRepoImpl) Delete(id string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM product WHERE id_product=?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}
func (s *ProductRepoImpl) Update(id string, product *models.Product) (*models.Product, error) {
	tx, err := s.db.Begin()
	_, err = tx.Exec("UPDATE product SET product_name=?,qty=?,type=?,warehouse_id=? WHERE id_product=?", &product.ProductName, &product.Qty, &product.Type, &product.ProductWarehouse.IdWarehouse, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return product, tx.Commit()
}
func InitProductRepoImpl(db *sql.DB) ProductRepository {
	return &ProductRepoImpl{db}
}
