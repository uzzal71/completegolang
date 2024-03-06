package main

import "fmt"

func passByValue(x int) {
	x = 10
}

func passByReference(x *int) {
    *x = 10
}

func main() {
	// pass by value
	value := 5
	passByValue(value)
    fmt.Println(value)

	// pass by reference
	pointer := 5
    passByReference(&pointer)
    fmt.Println(pointer)
}