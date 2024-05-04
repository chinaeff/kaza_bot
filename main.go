package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"kaza_bot/admin_bot"
	"kaza_bot/bot"
	"kaza_bot/database"
	"log"
)

func main() {
	dbProducts, err := sql.Open("sqlite3", "./database/products.db")
	if err != nil {
		log.Println("error opening database of products", err)
	}
	defer dbProducts.Close()

	dbOrders, err := sql.Open("sqlite3", "./database/orders.db")
	if err != nil {
		log.Println("error opening database of orders", err)
	}
	defer dbOrders.Close()

	dbUsers, err := sql.Open("sqlite3", "./database/users.db")
	if err != nil {
		log.Println("error opening database of users", err)
	}
	defer dbUsers.Close()

	err = database.CreateProductsTable(dbProducts)
	if err != nil {
		log.Println("error creating database of products", err)
	}
	err = database.CreateOrdersTable(dbOrders)
	if err != nil {
		log.Println("error creating database of orders", err)
	}
	err = database.CreateUsersTable(dbOrders)
	if err != nil {
		log.Println("error creating database of orders", err)
	}

	bot.StartBot()
	admin_bot.StartAdminBot()
}
