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
	} else {
		fmt.Println("C is false")
	}

	if a == 5 {
        fmt.Println("A is 5")
    } else {
        fmt.Println("A is not 5")
    }

	d := 25
	if d > 10 {
        fmt.Println("D is greater than 10")
    } else {
        fmt.Println("D is not greater than 10")
    }

	e := 20
	if e > 10 {
        fmt.Println("E is greater than 10")
    } else {
        fmt.Println("E is not greater than 10")
    }

	i := 14
}