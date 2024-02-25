package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}

	// accessing slice
	fmt.Println(mySlice[2])
	mySlice[2] = 100
	fmt.Println(mySlice[2])

	// appending slice
	mySlice = append(mySlice, 60, 70, 80)
    fmt.Println(mySlice)

    // iterating over slice
    for i, v := range mySlice {
        fmt.Printf("index: %d, value: %d\n", i, v)
    }

	// Slice of Slice 
	newSlice := mySlice[2:6]
	fmt.Println(newSlice)
}