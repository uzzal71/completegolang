package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}


func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Smith"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
}