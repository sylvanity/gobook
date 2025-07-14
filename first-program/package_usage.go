package main

import (
	"fmt"

	"example.com/first-program/mathutil"
)

func main() {
	sum := mathutil.Add(5, 10)
	fmt.Println("Sum from mathutil:", sum)
}
