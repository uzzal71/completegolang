package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type product struct {
	name string
	slug string
}


func main() {
	alex := person{firstName: "Uzzal",lastName: "Roy"}
	fmt.Println(alex.firstName, alex.lastName)
}