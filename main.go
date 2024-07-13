package main

import (
	"log"
	"net/http"

	"github.com/cinarci/myserver/handlers"
	"github.com/cinarci/myserver/models"
)

func main() {
	// Veritabanı bağlantısını başlat
	models.ConnectDatabase()

	// HTTP handler'ları ve middleware'i ayarla
	http.HandleFunc("/addresses", handlers.GetAddresses)
	http.HandleFunc("/address", handlers.CreateAddress)
	http.HandleFunc("/shipments", handlers.GetShipments)
	http.HandleFunc("/shipment", handlers.CreateShipment)
	http.HandleFunc("/generate-api-key", handlers.GenerateApiKey)

	// Middleware'i tüm rotalar için uygula
	http.Handle("/", handlers.ApiKeyMiddleware(http.DefaultServeMux))

	// Sunucuyu başlat
	log.Println("Sunucu 8080 portunda başlatıldı.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("sunucu başlatılamadı: %s\n", err)
	}
}
