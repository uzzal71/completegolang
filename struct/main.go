package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Uzzal","Roy"}
	fmt.Println(alex)
}