package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a context that can be cancelled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a channel to listen for OS signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle the signals.
	go func() {
		<-signalChan
		fmt.Println("\nReceived interrupt signal, shutting down gracefully...")
		cancel() // Cancel the context when a signal is received.
	}()

	// Subcommand definition
	nowCmd := flag.NewFlagSet("now", flag.ExitOnError)
	// forecastCmd is defined but not used in the book's snippet.
	// To avoid a "declared and not used" error, we can assign it to the blank identifier.
	_ = flag.NewFlagSet("forecast", flag.ExitOnError)

	// Check for subcommand
	if len(os.Args) < 2 {
		log.Fatal("Expected 'now' or 'forecast' subcommands")
	}

	// Get API Key
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENWEATHER_API_KEY environment variable not set")
	}
	client := NewAPIClient(apiKey)

	switch os.Args[1] {
	case "now":
		handleNowCommand(nowCmd, client, ctx)
	case "forecast":
		// handleForecastCommand(forecastCmd, client)
	default:
		log.Fatal("Unknown subcommand. Expected 'now' or 'forecast'")
	}
}

func handleNowCommand(cmd *flag.FlagSet, client *APIClient, ctx context.Context) {
	// Define flags specific to the 'now' command
	format := cmd.String("format", "text", "Output format: text or json")
	cmd.Parse(os.Args[2:])

	if cmd.NArg() == 0 {
		log.Fatal("City must be provided for 'now' command")
	}
	city := cmd.Arg(0)

	data, err := client.GetCurrentWeather(ctx, city)
	if err != nil {
		log.Fatalf("Could not get current weather: %v", err)
	}

	if *format == "json" {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal data to JSON: %v", err)
		}
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("Current weather in %s:\n", data.Name)
		fmt.Printf("  Temperature: %.1f°C\n", data.Main.Temp)
		if len(data.Weather) > 0 {
			fmt.Printf("  Conditions:  %s\n", data.Weather[0].Description)
		}
	}
}
