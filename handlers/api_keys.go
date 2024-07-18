package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

func GenerateApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := generateRandomApiKey(32) // 32 byte = 64 karakter hex string
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(apiKey))
}

func generateRandomApiKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
