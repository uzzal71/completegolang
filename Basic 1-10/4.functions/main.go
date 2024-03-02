package main

import "fmt"

func Sum() string {
	return "Hello World"
}

func AddSum() int {
	return 5 + 5
}

func main() {
	a := Sum()
	b := AddSum()
	fmt.Println(a, b)
}