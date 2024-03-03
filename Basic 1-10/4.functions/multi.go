package main

import "fmt"

type Person struct {
	Name string
	Age int
	Email string
	Phone string
	Address string
	Password string
	OTP string
}

func main() {
	num1 := 5
	num2 := 10
	result := num1 + num2
	fmt.Printf("%d\n", result)
}