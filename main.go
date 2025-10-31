package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ResolveRequest struct {
	Resolve []string `json:"resolve"`
}

type ResolveResponse struct {
	ResolvedValues map[string]string `json:"resolvedValues"`
}

func resolveValuesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ResolveRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Resolve environment variables
	resolvedValues := make(map[string]string)
	for _, envVarName := range req.Resolve {
		value := os.Getenv(envVarName)
		resolvedValues[envVarName] = value
		log.Printf("Resolved %s = %s", envVarName, value)
	}

	response := ResolveResponse{
		ResolvedValues: resolvedValues,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {
	http.HandleFunc("/test/resolve/values", resolveValuesHandler)
	http.HandleFunc("/test/resolve/values/", resolveValuesHandler)

	port := ":9090"
	log.Printf("Started service on port: %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
