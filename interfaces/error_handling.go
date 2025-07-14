package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// --- Basic Error Handling ---
func basicErrorHandling() {
	fmt.Println("--- Basic Error Handling ---")
	val, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("The converted value is:", val)

	val, err = strconv.Atoi("not-a-number")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}

// --- Creating Errors ---
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func openFile(name string) error {
	if name == "secret.txt" {
		return fmt.Errorf("access denied to file %q", name)
	}
	return nil
}

func creatingErrors() {
	fmt.Println("\n--- Creating Errors ---")
	_, err := divide(1, 0)
	fmt.Println("Error from errors.New:", err)

	err = openFile("secret.txt")
	fmt.Println("Error from fmt.Errorf:", err)
}

// --- Custom Error Types and Wrapping ---
type NetworkError struct {
	Timestamp time.Time
	URL       string
	Status    int
	Message   string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("at %v, request to %s failed with status %d: %s",
		e.Timestamp, e.URL, e.Status, e.Message)
}

func doRequest(url string) error {
	if url == "http://bad.example.com" {
		return &NetworkError{
			Timestamp: time.Now(),
			URL:       url,
			Status:    503,
			Message:   "service unavailable",
		}
	}
	return nil
}

// Fictional function that reads a config and wraps an error.
func readConfig() error {
	err := doRequest("http://bad.example.com")
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	return nil
}

func customErrorsAndWrapping() {
	fmt.Println("\n--- Custom Errors, Wrapping, errors.Is/As ---")
	err := readConfig()
	fmt.Println("Wrapped error:", err)

	// Check if the error is a specific type
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		fmt.Println("Assertion with errors.As successful.")
		fmt.Printf("Status: %d, Message: %s\n", netErr.Status, netErr.Message)
		if netErr.Status == 503 {
			fmt.Println("This is a temporary issue. We could retry.")
		}
	}

	// errors.Is can check for the original error inside the wrap
	originalError := doRequest("http://bad.example.com")
	wrappedErr := fmt.Errorf("another layer: %w", originalError)
	fmt.Println("Another wrapped error:", wrappedErr) // Use the variable
	// Note: errors.Is compares by value, so this check will fail as we create a new error instance.
	// For this to work, we'd need a sentinel error value.
	sentinel := errors.New("a specific error")
	errWithSentinel := fmt.Errorf("wrapped: %w", sentinel)
	if errors.Is(errWithSentinel, sentinel) {
		fmt.Println("errors.Is found the sentinel error in the chain.")
	}

}

func main() {
	basicErrorHandling()
	creatingErrors()
	customErrorsAndWrapping()
}
