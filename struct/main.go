package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}


func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Smith"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
}