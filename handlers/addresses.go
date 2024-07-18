package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cinarci/myserver/models"
)

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := models.GetAddresses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(addresses)
}

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateAddress(address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
