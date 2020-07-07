package services

import (
	"encoding/json"
	"logisticApi/main/master/middleware"
	"logisticApi/main/master/models"
	"logisticApi/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type WarehouseHandler struct {
	warehouseUsecase usecases.WarehouseUsecase
}

func WarehouseServices(r *mux.Router, service usecases.WarehouseUsecase) {
	WarehouseHandler := WarehouseHandler{service}
	r.Use(middleware.ActivityLogMiddleware)
	warehouse := r.PathPrefix("/warehouses").Subrouter()
	warehouse.HandleFunc("", WarehouseHandler.ListAllWarehouse).Methods(http.MethodGet)
	warehouse.HandleFunc("/{id}", WarehouseHandler.WarehouseId).Methods(http.MethodGet)
	warehouse.HandleFunc("/add", WarehouseHandler.addWarehouse).Methods(http.MethodPost)
	warehouse.HandleFunc("/update/{id}", WarehouseHandler.updateWarehouse).Methods(http.MethodPut)
	warehouse.HandleFunc("/delete/{id}", WarehouseHandler.deleteWarehouse).Methods(http.MethodDelete)
}
func (wh *WarehouseHandler) ListAllWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouses, err := wh.warehouseUsecase.GetAll()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	statusRespon := models.Responses{"Data Berhasil di ambil", http.StatusOK, warehouses}
	byteOfWarehouses, err := json.Marshal(statusRespon)
	if err != nil {
		w.Write([]byte("Oops,something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfWarehouses)
}
func (wh *WarehouseHandler) WarehouseId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	warehouse, err := wh.warehouseUsecase.GetById(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	statusRespon := models.Responses{"Data Berhasil di ambil", http.StatusOK, warehouse}
	byteOfWarehouse, err := json.Marshal(statusRespon)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteOfWarehouse)
}

func (wh *WarehouseHandler) addWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouseRequest []*models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&warehouseRequest)
	_, err := wh.warehouseUsecase.AddWarehouse(warehouseRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	byteData, err := json.Marshal(warehouseRequest)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (wh *WarehouseHandler) deleteWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := wh.warehouseUsecase.DeleteWarehouse(id)
	if err != nil {
		w.Write([]byte("Delete Data Failed!"))
	}
	var response models.Responses
	response.Status = http.StatusOK
	response.Messages = "Success Deleted Data"
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (wh *WarehouseHandler) updateWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&data)

	warehouse, err := wh.warehouseUsecase.UpdateWarehouse(id, &data)
	if err != nil {
		w.Write([]byte("Update Data Failed!"))
	}
	var response models.Responses
	response.Status = http.StatusOK
	response.Messages = "Success"
	response.Data = warehouse
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
