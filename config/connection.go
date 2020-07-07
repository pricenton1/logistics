package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func readGoEnvVar(key string) string {
	err := godotenv.Load("../config/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Connection() (*sql.DB, error) {
	dbEngine := readGoEnvVar("dbEngine")
	dbUser := readGoEnvVar("dbUser")
	dbPass := readGoEnvVar("dbPass")
	dbHost := readGoEnvVar("dbHost")
	dbPort := readGoEnvVar("dbPort")
	dbName := readGoEnvVar("dbName")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open(dbEngine, dbSource)
	if err := db.Ping(); err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Database Connected")
	}
	return db, nil
}
