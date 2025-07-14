package main

import "fmt"

// Titled is an interface that requires a Title() method.
type Titled interface {
	Title() string
}

// Book is a concrete type.
type Book struct {
	BookTitle string
	Author    string
}

// Title makes Book satisfy the Titled interface.
func (b Book) Title() string {
	return b.BookTitle
}

// Movie is another concrete type.
type Movie struct {
	MovieTitle string
	Director   string
}

// Title makes Movie satisfy the Titled interface.
func (m Movie) Title() string {
	return m.MovieTitle
}

// PrintTitle is a polymorphic function that works with any Titled value.
func PrintTitle(t Titled) {
	fmt.Printf("Content Title: %s\n", t.Title())
}

func main() {
	fmt.Println("--- Implicit Interface Satisfaction ---")
	myBook := Book{BookTitle: "The Go Programming Language", Author: "Donovan & Kernighan"}
	myMovie := Movie{MovieTitle: "Arrival", Director: "Denis Villeneuve"}

	PrintTitle(myBook)
	PrintTitle(myMovie)

	fmt.Println("\n--- Storing different types in an interface slice ---")
	content := []Titled{
		myBook,
		myMovie,
		Book{BookTitle: "Sapiens", Author: "Yuval Noah Harari"},
	}

	for _, item := range content {
		PrintTitle(item)
	}
}
