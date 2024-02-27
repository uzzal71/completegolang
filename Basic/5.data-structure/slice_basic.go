/**
* Creating Slices
* Accessing Elements
* Appending Slices
* Iterating Elements
* Copying a slice
 */

package main

import "fmt"

func main() {
	// Create slices
	var mySlice []int
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Printf("%T\n", mySlice)

	// slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	newSlice := arr[1:4]
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Printf("%T\n", newSlice)

	// slice initialization
	// var mySlice2 = []int{10, 20, 30, 40, 50}
	mySlice2 := []int{100, 200, 300, 400, 500}

	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Printf("%T\n", mySlice2)

	// using make
	mySlice3 := make([]int, 5)
    fmt.Println(mySlice3)
    fmt.Println(len(mySlice3))
    fmt.Printf("%T\n", mySlice3)

    // append
    mySlice3 = append(mySlice3, 60, 70, 80)
    fmt.Println(mySlice3)
    fmt.Println(len(mySlice3))
    fmt.Printf("%T\n", mySlice3)
}