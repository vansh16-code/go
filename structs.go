package main

import (
	"fmt"

)

type Product struct {
	ID    int
	Name  string
	Price int
}

// METHOD
func (p *Product) UpdatePrice(newPrice int) {
	p.Price = newPrice
}


func(p Product) Print(){
	fmt.Println("Product:" , p.Name ,"-" , p.Price)
}

func structDemo() {
	product := Product{ID: 1, Name: "Laptop", Price: 50000}

	fmt.Println("Before:", product)

	product.UpdatePrice(60000)

	fmt.Println("After:", product)

	product.Print()
}
