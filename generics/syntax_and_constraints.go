package main

import "fmt"

// --- Basic Generic Function ---
func PrintSlice[T any](s []T) {
	fmt.Print("[")
	for i, v := range s {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

// --- Using a Standard Interface as a Constraint ---
type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", i)
}

func Stringify[T fmt.Stringer](s []T) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = v.String()
	}
	return result
}

// --- Custom Constraint with Union and ~ Tilde ---
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func SumNumbers[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	fmt.Println("--- Generic PrintSlice ---")
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})

	fmt.Println("\n--- Generic Function with Interface Constraint ---")
	stringerSlice := []fmt.Stringer{MyInt(1), MyInt(2)}
	stringified := Stringify(stringerSlice)
	fmt.Println("Stringify result:", stringified)

	fmt.Println("\n--- Generic Function with Custom Constraint ---")
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of ints:", SumNumbers(intSlice))

	floatSlice := []float64{1.1, 2.2, 3.3}
	fmt.Println("Sum of floats:", SumNumbers(floatSlice))

	myInts := []MyInt{10, 20, 30}
	// This works because MyInt's underlying type is int, which matches ~int.
	fmt.Println("Sum of MyInts:", SumNumbers(myInts))
}
