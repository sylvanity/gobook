package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("--- Simple GET Request ---")
	// Using the default client for a simple GET
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	// CRITICAL: Always close the response body.
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	fmt.Println("Body (first 100 bytes):", string(body[:100]))

	fmt.Println("\n--- POST Request with Custom Client and Headers ---")
	// Create a custom client with a timeout.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create the request body.
	reqBody := strings.NewReader(`{"name":"Go", "purpose":"networking"}`)

	// Create the request object.
	req, err := http.NewRequest("POST", "http://httpbin.org/post", reqBody)
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	// Set custom headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer my-secret-token")
	req.Header.Set("X-Custom-Header", "Golang-Example")

	// Execute the request.
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	fmt.Println("Body contains 'Golang-Example':", strings.Contains(string(body), "Golang-Example"))
}
