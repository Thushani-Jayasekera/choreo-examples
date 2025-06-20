package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	targetURL = "https://c4ed2dec-7e72-4195-9032-81e66dd4d70f-dev-internal.e1-us-east-azure.internal.choreoapis.dev/rrissues/internalservice/v1.0"
	interval  = 2 * time.Minute
)

func invokeEndpoint() {
	resp, err := http.Get(targetURL)
	if err != nil {
		log.Printf("Error invoking endpoint: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	log.Printf("Received response: Status=%s, Body=%s", resp.Status, string(body))
}

func main() {
	log.Println("Starting cron service...")
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run once immediately
	invokeEndpoint()

	for range ticker.C {
		invokeEndpoint()
	}
} 