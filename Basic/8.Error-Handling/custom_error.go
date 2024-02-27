package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg string
	Code int
}

func (e CustomError) Error() string {
    return fmt.Sprintf("Error with code %d: %s", e.Code, e.Msg)
}

func SuperDivide(a, b int) (int, error) {
	if b == 0 {
       // return 0, CustomError{Msg: "division by zero", Code: 1}
	   return 0, errors.New("division by zero")
    }
	
	if a % b != 0 {
		return 0, CustomError{Msg: "not divisible", Code: 2}
	}

	return a / b, nil
}

func main() {
	result, err := SuperDivide(4, 3)
	if (err != nil) {
		var cErr CustomError
		if errors.As(err, &cErr) {
            fmt.Println("Customer error occurred")
			fmt.Println("Code", cErr.Code, "Message:", cErr.Msg)
        } else {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println(result)
}