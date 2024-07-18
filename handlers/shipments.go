package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cinarci/myserver/models"
)

func GetShipments(w http.ResponseWriter, r *http.Request) {
	shipments, err := models.GetShipments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shipments)
}

func CreateShipment(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateShipment(shipment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
