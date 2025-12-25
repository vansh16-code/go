package main

import "fmt"


type Product struct {
	ID    int    `json:"id"` // struct tag : `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}


func updatePrice(p *Product, newPrice int) {
	p.Price = newPrice
}


func structDemo() {

	product := Product{
		ID:    1,
		Name:  "Laptop",
		Price: 50000,
	}

	fmt.Println("Before update:", product)

	// Update price using pointer
	updatePrice(&product, 60000)

	fmt.Println("After update:", product)
}
