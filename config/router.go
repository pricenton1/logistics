package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter() // membuat router mux
	return router
}

func RunServer(router *mux.Router) {
	server := readGoEnvVar("serverHost")
	port := readGoEnvVar("serverPort")
	fmt.Println("Running Port :" + port)
	err := http.ListenAndServe(server+":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
