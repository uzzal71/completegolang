package main

import "fmt"

func main() {
	q := [4]int{1, 2, 3, 4}
	for i, v := range q {
		fmt.Printf("%d: %d\n", i, v)
	}
}