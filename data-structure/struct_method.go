package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person)Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func (p *Person) CelebrateBirthday() {
	p.Age++
	fmt.Printf("Happy Birthday %s!, your are now %d years old\n", p.Name, p.Age)
}

func main() {
	p := Person{"Uzzal Roy", 25}
	// fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	p.Print()
	p.CelebrateBirthday()
	p.Print()
}