package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("======================== TODO LIST ========================")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	fmt.Println(str)
}