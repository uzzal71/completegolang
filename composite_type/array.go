package main

import "fmt"

func main() {
	var a [3]int
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}
}