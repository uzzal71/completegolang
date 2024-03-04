package main

import "fmt"

func main() {
	a := 5
	b := 5.5
	c := true
	result := a + 10

    fmt.Printf("A = %d\n", a)
	fmt.Printf("B = %f\n", b)
	fmt.Printf("C = %t\n", c)
	fmt.Printf("Result = %d\n", result)

	if c == true {
		fmt.Println("C is true")
	}
}