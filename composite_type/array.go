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

	symbol := [...]string{USD: "$", EUR: "E", GBP: "EUR", RMB: "R"}
	fmt.Println(symbol[USD])
}