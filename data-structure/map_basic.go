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
}