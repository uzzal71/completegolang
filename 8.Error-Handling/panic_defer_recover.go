package main

import "fmt"

func deferFunc() {
	fmt.Println("defer function")
	if r := recover(); r != nil {
		fmt.Println("Recovered in defer function", r)
	}
}

func main() {
	defer deferFunc()
	
	fmt.Println("Before panic")
	panic("something bad happened")
	fmt.Println("After panic")
}