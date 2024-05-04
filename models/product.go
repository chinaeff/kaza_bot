package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
}

var Products = []Product{
	{ID: 1, Name: "Product 1", Price: 10, Image: "https://example.com/photo1.jpg"},
	{ID: 2, Name: "Product 2", Price: 19, Image: "https://example.com/photo2.jpg"},
}

func GetProducts() []Product {
	return Products
}
func AddProducts(product Product) {
	Products = append(Products, product)
}
