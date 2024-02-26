package main

import "fmt"

func deferFunc() {
	fmt.Println("defer function")
}

func main() {
	defer deferFunc()
	
	fmt.Println("Before panic")
	panic("something bad happened")
	fmt.Println("After panic")
}