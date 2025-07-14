package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // The '-' tag means "ignore this field" for json encoding
}

func main() {
	fmt.Println("--- Struct Tags for JSON Marshaling ---")
	user := User{
		ID:       123,
		Username: "gopher",
		Password: "a-very-secret-password",
	}

	// Marshal the struct into JSON. The tags will be used here.
	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
