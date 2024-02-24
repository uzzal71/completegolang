package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type product struct {
	name string
	slug string
	price float64
	stock int
}


func main() {
	product := product{name: "test", slug: "test", price: 25.36, stock: 100}
	fmt.Println(product)
}