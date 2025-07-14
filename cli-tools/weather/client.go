// In client.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Structs for parsing the JSON responses
type CurrentWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// ... (structs for forecast data would be defined here) ...

// APIClient manages communication with the weather API.
type APIClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewAPIClient(apiKey string) *APIClient {
	return &APIClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *APIClient) GetCurrentWeather(ctx context.Context, city string) (*CurrentWeatherData, error) {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s",
		city, c.apiKey,
	)

	// Create a new request with the context.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var data CurrentWeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// ... (a GetForecast method would be implemented similarly) ...
