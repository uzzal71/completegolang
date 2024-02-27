package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file.WriteString("Hello world!")

	fmt.Println("File created successfully")

	file.Close()
}