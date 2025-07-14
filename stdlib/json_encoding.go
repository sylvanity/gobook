package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

// --- Structs for JSON examples ---
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`    // Omit if the string is empty
	Password string `json:"-"`                  // The '-' tag means "always ignore this field"
	IsActive bool   `json:"isActive,omitempty"` // Example with a bool
}

// --- Marshal/Unmarshal example ---
func marshalUnmarshalExample() {
	fmt.Println("--- Marshal/Unmarshal Example ---")
	user := User{ID: 1, Username: "gopher", IsActive: true, Password: "secret"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %v", err)
	}
	fmt.Println("Marshaled JSON:", string(jsonData))

	// Unmarshal
	jsonDataToUnmarshal := []byte(`{"id":101, "username":"jane.doe", "isActive":true}`)
	var unmarshaledUser User
	err = json.Unmarshal(jsonDataToUnmarshal, &unmarshaledUser)
	if err != nil {
		log.Fatalf("JSON unmarshalling failed: %v", err)
	}
	fmt.Printf("Unmarshaled User: %+v\n", unmarshaledUser)
}

// --- Encoder/Decoder with HTTP Handler ---
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqUser User
	// Use a Decoder for the request body
	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// In a real app, you would process the user...
	fmt.Printf("Received user from request: %+v\n", reqUser)

	// Create a response
	responseUser := User{
		ID:       reqUser.ID,
		Username: reqUser.Username,
		IsActive: true, // Pretend we activated the user
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Use an Encoder for the response body
	err = json.NewEncoder(w).Encode(responseUser)
	if err != nil {
		// This is harder to handle as headers are already written
		log.Printf("Failed to encode response: %v", err)
	}
}

func encoderDecoderExample() {
	fmt.Println("\n--- Encoder/Decoder with HTTP Handler Example ---")

	// Create a mock request
	reqBody := `{"id":55, "username":"test.user", "email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	// Call the handler
	userHandler(rr, req)

	// Check the response
	fmt.Println("Handler Response Status:", rr.Code)
	fmt.Println("Handler Response Body:", rr.Body.String())
}

func main() {
	marshalUnmarshalExample()
	encoderDecoderExample()
}
