package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// Configuration holds the environment variables for the service.
type Configuration struct {
	ServiceURL   string
	TokenURL     string
	ClientID     string
	ClientSecret string
}

// GreetingResponse represents the response from the hello service.
type GreetingResponse struct {
	Message string `json:"message"`
}

// DiagnosticInfo holds diagnostic information.
type DiagnosticInfo struct {
	ServiceURL        string `json:"serviceUrl"`
	TokenURL          string `json:"tokenUrl"`
	ClientID          string `json:"clientId"`
	ClientSecret      string `json:"clientSecret"`
	DiagnosticVersion string `json:"diagnosticVersion"`
}

// GetConfiguration retrieves the configuration from environment variables.
func GetConfiguration() (*Configuration, error) {
	return &Configuration{
		ServiceURL:   os.Getenv("SERVICE_URL"),
		TokenURL:     os.Getenv("TOKEN_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}, nil
}

// GreetingHandler returns a greeting message from an external service.
func GreetingHandler(config *Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subpath := strings.TrimPrefix(r.URL.Path, "/greeting")
		if subpath != "" {
			subpath = "/" + subpath
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", config.ServiceURL+subpath, nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to get greeting", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var result GreetingResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

// DiagnosticHandler returns diagnostic information about the service.
func DiagnosticHandler(config *Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := DiagnosticInfo{
			ServiceURL:        config.ServiceURL,
			TokenURL:          config.TokenURL,
			ClientID:          config.ClientID,
			ClientSecret:      config.ClientSecret,
			DiagnosticVersion: "v1.0",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	}
}

func main() {
	config, err := GetConfiguration()
	if err != nil {
		log.Fatalf("Failed to get configuration: %v", err)
	}

	http.HandleFunc("/greeting", GreetingHandler(config))
	http.HandleFunc("/diagnostic", DiagnosticHandler(config))

	log.Println("Starting server on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
