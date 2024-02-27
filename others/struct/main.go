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
	jim := person{
		firstName: "Uzzal",
		lastName: "Roy",
		contact: contactInfo{
			email: "uzzal@gmail.com",
			zipCode: 5300,
		},
	}

	fmt.Printf("%+v\n", jim)
}