/**
* Creating Structs
* Accessing Field
* Modifying Field
* Method on Struct
 */

package main

import "fmt"

func main() {
	// creating struct
	type Person struct {
		Name string
		Age int
	}

	var p Person
	fmt.Println(p)

	// Accessing values
	var pr Person = Person{"Uzzal", 25}
	fmt.Println(pr)
}