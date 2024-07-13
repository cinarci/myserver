package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cinarci/myserver/models"

	"github.com/google/uuid"
)

// Handler to handle GET requests for addresses
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	addresses, err := models.GetAddresses(userID)
	if err != nil {
		log.Printf("Error querying database: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

// Handler to handle POST requests for addresses
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = models.CreateAddress(address)
	if err != nil {
		log.Printf("Error inserting into database: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// Handler to handle GET requests for shipments
func GetShipments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	shipments, err := models.GetShipments(userID)
	if err != nil {
		log.Printf("Error querying database: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipments)
}

// Handler to handle POST requests for shipments
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = models.CreateShipment(shipment)
	if err != nil {
		log.Printf("Error inserting into database: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipment)
}

// Middleware to validate API key
func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if !models.ValidateAPIKey(apiKey) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Handler to generate API key
func GenerateApiKey(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	apiKey := uuid.New().String()
	err := models.SaveAPIKey(username, apiKey)
	if err != nil {
		log.Printf("Error saving API key: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Generated API key for %s: %s\n", username, apiKey)
}
