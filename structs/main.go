package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	first_name string
	last_name  string
	contact    contactInfo
}

func main() {
	sujon := person{
		first_name: "Sujon",
		last_name:  "Roy",
		contact: contactInfo{
			email:   "sujon@gmail.com",
			zipCode: 8700,
		},
	}
	tapos := person{
		first_name: "Tapos",
		last_name:  "Roy",
		contact: contactInfo{
			email:   "tapos@gmail.com",
			zipCode: 4500,
		},
	}

	var uzzal person

	uzzal.first_name = "Uzzal"
	uzzal.last_name = "Roy"
	uzzal.contact.email = "uzzal@gmail.com"
	uzzal.contact.zipCode = 5300

	fmt.Printf("%+v\n", sujon)
	fmt.Printf("%+v\n", tapos)
	fmt.Printf("%+v\n", uzzal)
}
