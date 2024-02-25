package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p := Person{"Uzzal Roy", 25}
	fmt.Println(p)

	// Accessing fields
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

	// Modifying fields
	p.Name = "Sujon Roy"
	p.Age = 30
    fmt.Println(p)
}