package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

// --- Simple Handlers with Default ServeMux ---

func homeHandlerSimple(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Welcome to the simple home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}

// --- Handler as a method on a struct ---

type Server struct {
	// In a real app, this would be a real DB connection.
	// We use nil here for demonstration purposes.
	db *sql.DB
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome from the structured server!")
	// This handler could now access s.db if it were initialized.
}

func main() {
	// Example 1: Using the default ServeMux
	// This server will run in a goroutine and be shut down.
	go func() {
		http.HandleFunc("/", homeHandlerSimple)
		http.HandleFunc("/about", aboutHandler)
		log.Println("Starting simple server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
			log.Fatalf("Simple server failed: %v", err)
		}
		log.Println("Simple server shut down.")
	}()

	// Example 2: Using a custom ServeMux and a server struct
	// This server will also run in a goroutine.
	go func() {
		server := &Server{db: nil}
		mux := http.NewServeMux()
		mux.HandleFunc("/", server.homeHandler)
		log.Println("Starting structured server on :8081")
		if err := http.ListenAndServe(":8081", mux); err != http.ErrServerClosed {
			log.Fatalf("Structured server failed: %v", err)
		}
		log.Println("Structured server shut down.")
	}()

	// Let the servers run for a moment, then we can imagine shutting them down.
	// In a real app, this would be handled by os.Signal.
	// For this example, we just wait a bit.
	log.Println("Servers running for 2 seconds...")
	time.Sleep(2 * time.Second)
	log.Println("Example finished. In a real app, you would need graceful shutdown.")
}
