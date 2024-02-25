/**
* Creating Maps
* Accessing Elements
* Modifying Slices
* Iterating over a map
* Deleting Elements
 */

package main

import "fmt"

func main() {
	// creating a map
	var myMap map[string]int
	fmt.Println(myMap)

	// initilization
	myMap2 := map[string]int{
		"apple": 5,
		"banana": 10,
        "cherry": 15,
	}
	fmt.Println(myMap2)

	// accessing elements
	fmt.Println(myMap2["apple"])
    myMap2["apple"] = 100
    fmt.Println(myMap2["apple"])

	// add new key
	myMap2["orange"] = 200
    fmt.Println(myMap2)

    // deleting elements
    delete(myMap2, "banana")
    fmt.Println(myMap2)

    // iterating over map
    for k, v := range myMap2 {
        fmt.Printf("key: %s, value: %d\n", k, v)
    }
}