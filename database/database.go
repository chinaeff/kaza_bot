package database

import (
	"database/sql"
)

func CreateProductsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    price INTEGER,
    image BLOB
)`)
	return err
}
func CreateOrdersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
    order_id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER,
    status TEXT,
    date_of_order DATETIME,
    user_id INTEGER,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)`)
	return err
}
func CreateUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
    users_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    orders_json TEXT
)`)
	return err
}
