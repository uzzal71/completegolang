package main

import "fmt"

func main() {
	var s []int = []int{1, 2, 3} // [3]int{1, 2, 3} // len=3 // capacity=3
	fmt.Println(s, &s[0])
	s = append(s, 4)  // [6]int{1, 2, 3, 4} // len=4 // capacity=6
	fmt.Println(s, &s[0])

	s = append(s, 5) // [6]int{1, 2, 3, 4} // len=4 // capacity=6
	fmt.Println(s, &s[0])
}