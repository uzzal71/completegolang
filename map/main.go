package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)
	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)

	colorAgains := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
		"white": "#ffffff",
	}

	printMap(colorAgains)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
