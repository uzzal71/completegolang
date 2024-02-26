package main

import "fmt"

func main() {
	// new function
	var p *int
	p = new(int)
	*p = 5
	fmt.Println(p, *p)

	// make function

}