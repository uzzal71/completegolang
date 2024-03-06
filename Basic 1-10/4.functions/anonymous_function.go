package main

import "fmt"


func main() {
	result := func(x int, y int) int{
		return x + y
	}(5, 6)
	fmt.Println(result)
}