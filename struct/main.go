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
}


func main() {
	alex := person{firstName: "Uzzal",lastName: "Roy"}
	fmt.Println(alex.firstName, alex.lastName)
}