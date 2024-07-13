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

// Adresler için get isteklerini işleyen handler
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "İzin verilmeyen Method", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Geçersiz User ID", http.StatusBadRequest)
		return
	}

	addresses, err := models.GetAddresses(userID)
	if err != nil {
		log.Printf("Veri tabanı sorgusunda hata: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

// Adresler için post isteklerini işleyen handler
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "İzin verilmeyen method", http.StatusMethodNotAllowed)
		return
	}

	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		log.Printf("İstek gövdesi işlenirken hata oluştu: %s", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = models.CreateAddress(address)
	if err != nil {
		log.Printf("Veri tabanına eklemede hata: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// Gönderiler için get isteklerini işleyen handler
func GetShipments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "İzin verilmeyen method", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	shipments, err := models.GetShipments(userID)
	if err != nil {
		log.Printf("Veri tabanı sorgusunda hata: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipments)
}

// Gönderi oluşturma için Post isteklerini işleyen handler
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "İzin verilmeyen method", http.StatusMethodNotAllowed)
		return
	}

	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		log.Printf("İstek gövdesi işlenirken hata oluştu: %s", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = models.CreateShipment(shipment)
	if err != nil {
		log.Printf("Veri tabanına eklemede hata: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipment)
}

// API anahtarını doğrulamak için middleware
func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if !models.ValidateAPIKey(apiKey) {
			http.Error(w, "Bu işlem için yetkiniz yok", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Api anahtarı oluşturmak için handler
func GenerateApiKey(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Kullanıcı adı gereklidir", http.StatusBadRequest)
		return
	}

	apiKey := uuid.New().String()
	err := models.SaveAPIKey(username, apiKey)
	if err != nil {
		log.Printf("API anahtarı kaydedilemedi.: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, username, "için oluşturulan API anahtarı %s: %s\n", apiKey)
}
