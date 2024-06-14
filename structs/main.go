package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	first_name string
	last_name  string
	contactInfo
}

func main() {
	sujon := person{
		first_name: "Sujon",
		last_name:  "Roy",
		contactInfo: contactInfo{
			email:   "sujon@gmail.com",
			zipCode: 8700,
		},
	}
	tapos := person{
		first_name: "Tapos",
		last_name:  "Roy",
		contactInfo: contactInfo{
			email:   "tapos@gmail.com",
			zipCode: 4500,
		},
	}

	var uzzal person

	uzzal.first_name = "Uzzal"
	uzzal.last_name = "Roy"
	uzzal.contactInfo.email = "uzzal@gmail.com"
	uzzal.contactInfo.zipCode = 5300

	uzzalPointer := &uzzal
	uzzalPointer.updateName("Uzzal Kumar")

	tapos.updateName("Tapos Kumar")

	sujon.print()
	tapos.print()
	uzzal.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).first_name = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
