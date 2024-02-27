package main

import "fmt"

func Multiply(a, b int, ch chan int) {
	ch <- a * b
}

func main() {
	ch := make(chan int) // Buffered channel with capacity 1

	go Multiply(2, 3, ch)

	res := <-ch
	fmt.Println(res)
}
