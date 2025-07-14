package main

import "fmt"

// Employee is a struct that groups together related employee information.
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Salary    float64
	IsActive  bool
}

// FullName is a method with a value receiver.
// It operates on a copy of the Employee.
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// GiveRaise is a method with a pointer receiver.
// It operates on the original Employee value and can modify it.
func (e *Employee) GiveRaise(percentage float64) {
	e.Salary += e.Salary * (percentage / 100.0)
}

func main() {
	fmt.Println("--- Structs and Methods ---")

	// Using a struct literal with named fields is the recommended way.
	jane := Employee{
		ID:        102,
		FirstName: "Jane",
		LastName:  "Doe",
		Salary:    75000.0,
		IsActive:  true,
	}
	fmt.Printf("Initial Employee: %+v\n", jane)

	// Calling a method with a value receiver.
	fmt.Println("Full Name:", jane.FullName())

	// Calling a method with a pointer receiver.
	// Go automatically converts 'jane' to '&jane'.
	jane.GiveRaise(10)
	fmt.Printf("Employee after 10%% raise: %+v\n", jane)

	// Although jane is a value, we called a pointer receiver method.
	// The original 'jane' variable was modified.
	fmt.Println("Salary after raise:", jane.Salary)
}
