package main

import (
	"fmt"
	"shapes/rectangle"
)

func main() {
	r := rectangle.Rectangle{2, 4}
	area := r.Area()
	fmt.Println(area)
}