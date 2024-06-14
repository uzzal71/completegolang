package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

func main() {
	sujon := person{"Sujon", "Roy"}
	tapos := person{first_name: "Tapos", last_name: "Roy"}

	fmt.Println(sujon)
	fmt.Println(tapos)
}
