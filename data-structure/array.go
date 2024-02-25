/**
* Declaration
* Initialization
* Accessing Elements
* Modifying Elements
* Iterating over an Array OR Looping
 */

package main

import "fmt"

func main() {
	// ************** (declaration) ******************
	var arrayName [5]int 
	fmt.Println(arrayName)
	fmt.Println(len(arrayName))
	fmt.Printf("%T\n", arrayName)

	// **************** (initialization) ****************
	// var arrId[5]int = [5]int{0, 1, 2, 3} 
	// var arrId = [5]int{0, 1, 2, 3}
	// arrId := [5]int{0, 1, 2, 3} 

	arrId := [...]int{0, 1, 2, 3} 

	fmt.Println(arrId)
	fmt.Println(len(arrId))
	fmt.Printf("%T\n", arrId)
}