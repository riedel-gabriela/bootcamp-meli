package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationMigrate{
		Db: &mysql.Config{
			User:   "root",
			Passwd: "",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		FilePathCustomer: "./docs/db/json/customers.json",
		FilePathProduct:  "./docs/db/json/products.json",
		FilePathInvoice:  "./docs/db/json/invoices.json",
		FilePathSale:     "./docs/db/json/sales.json",
	}
	app := application.NewApplicationMigrate(cfg)
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
