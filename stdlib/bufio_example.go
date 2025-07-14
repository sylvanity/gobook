package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a temporary file to use for the example
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("temp.txt") // Clean up the file

	// --- bufio.Writer example ---
	writer := bufio.NewWriter(file)
	linesWritten := []string{
		"Line 1: The quick brown fox",
		"Line 2: jumps over the lazy dog.",
		"Line 3: The end.",
	}
	for _, line := range linesWritten {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("failed to write string: %v", err)
		}
	}
	// It's crucial to Flush to ensure all data is written to the underlying writer.
	writer.Flush()
	file.Close() // Close the file to save it.

	// --- bufio.Reader example ---
	file, err = os.Open("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("--- Reading file with bufio.Reader ---")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read line: %v", err)
		}
	}
	fmt.Println("\n--- Finished reading ---")
}
