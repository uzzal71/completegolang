package main

import "fmt"


func main() {
	result := func(x int, y int) int{
		return x + y
	}(5, 6)
	fmt.Println(result)

	sum := func (x int, y int) int {
		return x + y
	}

	result2 := sum(2, 3)
	fmt.Println(result2)

	a := 10
	fmt.Println(a)
}