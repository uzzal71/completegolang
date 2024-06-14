package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	first_name string
	last_name  string
	age        int
	contactInfo
}

func main() {
	// uzzal := person{"Uzzal", "Roy", 25}
	// fmt.Println(uzzal)

	sujon := person{
		first_name: "Sujon",
		last_name:  "Roy",
		age:        28,
		contactInfo: contactInfo{
			email:   "sujon@gmail.com",
			zipCode: 8700,
		},
	}
	sujon.updateName("Sujon Kumar")
	sujon.print()

	var tapos person
	tapos.first_name = "Tapos"
	tapos.last_name = "Roy"
	tapos.age = 27
	tapos.contactInfo.email = "tapos@gmail.com"
	tapos.contactInfo.zipCode = 4500
	tapos.updateName("Tapos Kumar")
	tapos.print()
}

func (p person) updateName(newFirstName string) {
	p.first_name = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
