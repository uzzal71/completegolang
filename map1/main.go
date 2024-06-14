package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := map[string]string{
	// 	"white": "#ffffff",
	// 	"black": "#000000",
	// }

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	varients := map[int]string{
		10: "Red",
		20: "Green",
		30: "Yellow",
	}

	fmt.Println(varients)

	colors := map[string]string{
		"gray":  "#f1f1f1",
		"white": "#ffffff",
		"black": "#000000",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
