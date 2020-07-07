package master

import (
	"database/sql"
	"logisticApi/main/master/repositories"
	"logisticApi/main/master/services"
	"logisticApi/main/master/usecases"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	// product init
	productRepo := repositories.InitProductRepoImpl(db)
	productUsecase := usecases.InitProductUsecaseImpl(productRepo)
	services.ProductServices(r, productUsecase)

}
