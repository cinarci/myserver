package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB global database variable
var DB *sql.DB

// User struct for user data
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	ApiKey    string `json:"api_key"`
	CreatedAt string `json:"created_at"`
}

// Address struct for address data
type Address struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	CreatedAt  string `json:"created_at"`
}

// Shipment struct for shipment data
type Shipment struct {
	ID                int     `json:"id"`
	UserID            int     `json:"user_id"`
	SenderAddressID   int     `json:"sender_address_id"`
	ReceiverAddressID int     `json:"receiver_address_id"`
	Weight            float64 `json:"weight"`
	Status            string  `json:"status"`
	CreatedAt         string  `json:"created_at"`
}

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	var err error

	// Replace these parameters with your database configuration
	dsn := "root:@tcp(127.0.0.1:3306)/cargo"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	fmt.Println("Database connection established")
}

// SaveAPIKey saves the generated API key to the database
func SaveAPIKey(username, apiKey string) error {
	_, err := DB.Exec("INSERT INTO users (username, api_key) VALUES (?, ?)", username, apiKey)
	return err
}

// ValidateAPIKey checks if the provided API key is valid
func ValidateAPIKey(apiKey string) bool {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE api_key = ?", apiKey).Scan(&count)
	if err != nil {
		log.Printf("Error querying database: %s", err)
		return false
	}
	return count > 0
}

// CreateAddress saves a new address to the database
func CreateAddress(address Address) error {
	_, err := DB.Exec("INSERT INTO addresses (user_id, address, city, postal_code) VALUES (?, ?, ?, ?)",
		address.UserID, address.Address, address.City, address.PostalCode)
	return err
}

// GetAddresses retrieves addresses for a user
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

// CreateShipment saves a new shipment to the database
func CreateShipment(shipment Shipment) error {
	_, err := DB.Exec("INSERT INTO shipments (user_id, sender_address_id, receiver_address_id, weight, status) VALUES (?, ?, ?, ?, ?)",
		shipment.UserID, shipment.SenderAddressID, shipment.ReceiverAddressID, shipment.Weight, shipment.Status)
	return err
}

// GetShipments retrieves shipments for a user
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
