package main

import "fmt"

type CustomError struct {
	Msg string
	Code int
}

func (e CustomError) Error() string {
    return fmt.Sprintf("Error with code %d: %s", e.Code, e.Msg)
}

func SuperDivide(a, b int) (int, error) {
	if b == 0 {
        return 0, CustomError{Msg: "division by zero", Code: 1}
    }
	
	if a % b != 0 {
		return 0, CustomError{Msg: "not divisible", Code: 2}
	}

	return a / b, nil
}

func main() {
	result, err := SuperDivide(4, 0)
	if (err != nil) {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println(result)
}