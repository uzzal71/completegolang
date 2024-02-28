package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("======================== TODO LIST ========================")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Println("============================================================")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()
		fmt.Println(choice)
	}
}