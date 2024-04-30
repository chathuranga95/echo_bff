package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Config struct for holding environment variables
type Config struct {
	ServiceURL string
}

// Get the service URL from the environment
func getServiceURL() string {
	return os.Getenv("SERVICE_URL")
}

// GreetingHandler for /greeting endpoint
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	serviceURL := getServiceURL()

	// Create an HTTP client to fetch the greeting
	client := &http.Client{}
	log.Println("Fetching greeting from", serviceURL+"/"+r.URL.Query().Get("subpath"))
	resp, err := client.Get(serviceURL + "/" + r.URL.Query().Get("subpath"))
	if err != nil {
		http.Error(w, "Failed to fetch greeting", http.StatusInternalServerError)
		log.Println("Error fetching greeting:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response and write it back as JSON
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		log.Println("Error reading response:", err)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		log.Println("Error writing response:", err)
	}
	log.Println("Greeting response sent")
}

// DiagnosticHandler for /diagnostic endpoint
func DiagnosticHandler(w http.ResponseWriter, r *http.Request) {
	diagnostic := map[string]interface{}{
		"serviceUrl":        getServiceURL(),
		"diagnosticVersion": "v1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diagnostic)

	log.Println("Diagnostic response sent")
}

func main() {
	http.HandleFunc("/greeting", GreetingHandler)
	http.HandleFunc("/diagnostic", DiagnosticHandler)

	port := ":9090"
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
