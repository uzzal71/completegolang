package main

import "fmt"

func printNumbers() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
        fmt.Printf("%c\n", i)
    }
}

func main() {
	
}