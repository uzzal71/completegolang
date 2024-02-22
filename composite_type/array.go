package main

import "fmt"

func main() {
	var q [3]int = [3]int{1, 2, 3}
	var r [2]int = [2]int{1, 2}

	for i, v := range q {
		fmt.Printf("%d: %d\n", i, v)
	}
	for i, v := range r {
        fmt.Printf("%d: %d\n", i, v)
    }
}