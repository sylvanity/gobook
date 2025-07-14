package main

import (
	"errors"
	"fmt"
)

// add demonstrates basic function definition.
func add(x, y int) int {
	return x + y
}

// divide demonstrates returning multiple values, including an error.
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil // 'nil' is the zero value for error types, indicating success.
}

// divideNamed demonstrates named return values.
func divideNamed(x, y float64) (result float64, err error) {
	if y == 0 {
		err = errors.New("cannot divide by zero")
		return // Returns 0 and the error
	}
	result = x / y
	return // Returns the result and a nil error
}

func main() {
	fmt.Println("--- add ---")
	sum := add(40, 2)
	fmt.Println("40 + 2 =", sum)

	fmt.Println("\n--- divide ---")
	result1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("10 / 2 =", result1)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("10 / 0 =", result2)
	}

	fmt.Println("\n--- divideNamed ---")
	result3, err3 := divideNamed(10, 2)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("10 / 2 =", result3)
	}
}
