package main

import (
	"logisticApi/config"
	"logisticApi/main/master"
)

func main() {
	db, _ := config.Connection()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
