package main

import (
	"fmt"
	"time"
)

func Multiply(a, b int, ch chan int) {
	time.Sleep(250 * time.Microsecond)
	ch <- a * b
}

func main() {
	ch := make(chan int) // Buffered channel with capacity 1

	go Multiply(2, 3, ch)

	res := <-ch
	fmt.Println(res)

	close(ch)

	res = <-ch
	fmt.Println(res)
}
