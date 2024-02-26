package main

import "fmt"

func divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println(divide(4, 0))
}