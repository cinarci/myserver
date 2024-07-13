package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Kullanıcı veriler için user struct'ı
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	ApiKey    string `json:"api_key"`
	CreatedAt string `json:"created_at"`
}

// Adres verileri için adress struct'ı
type Address struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	CreatedAt  string `json:"created_at"`
}

// Gönderi verileri için gönderi struct'ı
type Shipment struct {
	ID                int     `json:"id"`
	UserID            int     `json:"user_id"`
	SenderAddressID   int     `json:"sender_address_id"`
	ReceiverAddressID int     `json:"receiver_address_id"`
	Weight            float64 `json:"weight"`
	Status            string  `json:"status"`
	CreatedAt         string  `json:"created_at"`
}

// Veri tabanı bağlantısını başlatan metod
func ConnectDatabase() {
	var err error

	// lokalde çalışan mysql veritabanı için bağlantı bilgisi.
	dsn := "root:@tcp(127.0.0.1:3306)/cargo"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Veri tabanı açılırken hata oluştu: %s", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Veritabanı bağlantısında hata: %s", err)
	}

	fmt.Println("Veri tabanı bağlantısı yapılmış durumda")
}

// Oluşturulan API anahtarını veri tabanına kaydeden metod
func SaveAPIKey(username, apiKey string) error {
	_, err := DB.Exec("INSERT INTO users (username, api_key) VALUES (?, ?)", username, apiKey)
	return err
}

// Sağlanan API anahtarını doğrulayan metod
func ValidateAPIKey(apiKey string) bool {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE api_key = ?", apiKey).Scan(&count)
	if err != nil {
		log.Printf("Error querying database: %s", err)
		return false
	}
	return count > 0
}

// Veri tabanında yeni bir adress oluşturma
func CreateAddress(address Address) error {
	_, err := DB.Exec("INSERT INTO addresses (user_id, address, city, postal_code) VALUES (?, ?, ?, ?)",
		address.UserID, address.Address, address.City, address.PostalCode)
	return err
}

// Bir kullanıcının adresini alma
func GetAddresses(userID int) ([]Address, error) {
	rows, err := DB.Query("SELECT id, user_id, address, city, postal_code, created_at FROM addresses WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []Address
	for rows.Next() {
		var address Address
		if err := rows.Scan(&address.ID, &address.UserID, &address.Address, &address.City, &address.PostalCode, &address.CreatedAt); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

// Veritabanında yeni bir gönderi oluşturma
func CreateShipment(shipment Shipment) error {
	_, err := DB.Exec("INSERT INTO shipments (user_id, sender_address_id, receiver_address_id, weight, status) VALUES (?, ?, ?, ?, ?)",
		shipment.UserID, shipment.SenderAddressID, shipment.ReceiverAddressID, shipment.Weight, shipment.Status)
	return err
}

// Bir kullanıcı için adresleri getirir.
func GetShipments(userID int) ([]Shipment, error) {
	rows, err := DB.Query("SELECT id, user_id, sender_address_id, receiver_address_id, weight, status, created_at FROM shipments WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shipments []Shipment
	for rows.Next() {
		var shipment Shipment
		if err := rows.Scan(&shipment.ID, &shipment.UserID, &shipment.SenderAddressID, &shipment.ReceiverAddressID, &shipment.Weight, &shipment.Status, &shipment.CreatedAt); err != nil {
			return nil, err
		}
		shipments = append(shipments, shipment)
	}
	return shipments, nil
}
