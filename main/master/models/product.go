package models

type Product struct {
	IdProduct        string `json:"id"`
	ProductName      string `json:"productName"`
	Qty              string `json:"quantity"`
	Type             string `json:"typeProduct"`
	ProductWarehouse Warehouse
}
