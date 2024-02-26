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

	// copying slice
	newSlice2 := make([]int, len(newSlice))
    copy(newSlice2, newSlice)
    fmt.Println(newSlice2)

	newSlice2[2] = 500;
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
}