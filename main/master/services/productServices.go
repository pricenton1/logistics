package services

import (
	"encoding/json"
	"logisticApi/main/master/middleware"
	"logisticApi/main/master/models"
	"logisticApi/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUsecase usecases.ProductUsecase
}

func ProductServices(r *mux.Router, service usecases.ProductUsecase) {
	ProductHandler := ProductHandler{service}
	r.Use(middleware.ActivityLogMiddleware)

	product := r.PathPrefix("/products").Subrouter()
	product.HandleFunc("", ProductHandler.ListAllProduct).Methods(http.MethodGet)
	product.HandleFunc("/{id}", ProductHandler.ProductId).Methods(http.MethodGet)
	product.HandleFunc("/add", ProductHandler.addProduct).Methods(http.MethodPost)
	product.HandleFunc("/update{id}", ProductHandler.updateProduct).Methods(http.MethodPut)
	product.HandleFunc("/delete{id}", ProductHandler.deleteProduct).Methods(http.MethodDelete)
}
func (p *ProductHandler) ListAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := p.productUsecase.GetAll()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	statusRespon := models.Responses{"Data Berhasil di ambil", http.StatusOK, products}
	byteOfproducts, err := json.Marshal(statusRespon)
	if err != nil {
		w.Write([]byte("Oops,something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfproducts)
}
func (p *ProductHandler) ProductId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	product, err := p.productUsecase.GetById(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	statusRespon := models.Responses{"Data Berhasil di ambil", http.StatusOK, product}
	byteOfProduct, err := json.Marshal(statusRespon)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteOfProduct)
}

func (p *ProductHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest []*models.Product
	_ = json.NewDecoder(r.Body).Decode(&productRequest)
	_, err := p.productUsecase.AddProduct(productRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	byteData, err := json.Marshal(productRequest)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (p *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := p.productUsecase.DeleteProduct(id)
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
func (p *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Product
	_ = json.NewDecoder(r.Body).Decode(&data)

	product, err := p.productUsecase.UpdateProduct(id, &data)
	if err != nil {
		w.Write([]byte("Update Data Failed!"))
	}
	var response models.Responses
	response.Status = http.StatusOK
	response.Messages = "Success"
	response.Data = product
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
