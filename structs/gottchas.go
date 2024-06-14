package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "Uzzal Kumar Roy"
	fmt.Println(*&name)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
