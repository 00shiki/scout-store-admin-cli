package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"scout-store-admin-cli/cli"
	"scout-store-admin-cli/config"
	"scout-store-admin-cli/handler"
	PRODUCT_REPO "scout-store-admin-cli/repository/product"
	SALES_REPO "scout-store-admin-cli/repository/sales"
	STAFF_REPO "scout-store-admin-cli/repository/staff"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	db, err := sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	productRepo := PRODUCT_REPO.NewRepository(db)
	staffRepo := STAFF_REPO.NewRepository(db)
	salesRepo := SALES_REPO.NewRepository(db)

	h := handler.NewHandler(productRepo, staffRepo, salesRepo)
	c := cli.New(h)
	c.Init()
}
