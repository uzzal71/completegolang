package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "you"}
	upddateSlice(mySlice)
	fmt.Println(mySlice)
}

func upddateSlice(s []string) {
	s[0] = "Bye"
}
