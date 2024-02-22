package main

import "fmt"

func main() {
	type Currency int

	const (
        USD Currency = iota
        EUR
        GBP
		RMB
    )

	symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}
	
	for i, v := range symbol {
		fmt.Println(i, v)
	}
}