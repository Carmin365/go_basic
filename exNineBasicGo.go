package main

import (
	"errors"
	"fmt"
)

func division(dividend int, divider int) (result int, failure error) {
	if divider == 0 {
		failure = errors.New("Division by zero!")
		return
	}

	result = dividend / divider
	return
}

func main() {
	result, failure := division(10, 2)

	if failure != nil {
		fmt.Println("Failure:", failure)
	} else {
		fmt.Println("Result:", result)
	}
}
