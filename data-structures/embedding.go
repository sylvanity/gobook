package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type ContactInfo struct {
	Email string
	Phone string
}

// Manager embeds only Employee, matching the text example.
type Manager struct {
	Employee   // Embedded anonymous field
	Department string
}

func main() {
	fmt.Println("--- Struct Embedding (Composition) ---")
	mgr := Manager{
		Employee: Employee{
			FirstName: "John",
			LastName:  "Smith",
		},
		Department: "Engineering",
	}

	// We can access the embedded fields directly.
	fmt.Println("Manager's First Name:", mgr.FirstName)
	fmt.Println("Manager's Last Name:", mgr.LastName)
	fmt.Println("Manager's Department:", mgr.Department)

	// The full struct literal shows the structure.
	fmt.Printf("Full Manager struct: %+v\n", mgr)
}
