package main

import "fmt"

/*
func calculate(x int, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}
*/

func calculate(x int, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return
}

func main() {
	s, d := calculate(10, 5)
	fmt.Println(s, d)
}