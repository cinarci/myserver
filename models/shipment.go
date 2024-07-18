package models

type Shipment struct {
	ID         uint `gorm:"primaryKey"`
	AddressID  uint
	TrackingID string
	Status     string
}

func GetShipments() ([]Shipment, error) {
	var shipments []Shipment
	if err := db.Find(&shipments).Error; err != nil {
		return nil, err
	}
	return shipments, nil
}

func CreateShipment(shipment Shipment) error {
	if err := db.Create(&shipment).Error; err != nil {
		return err
	}
	return nil
}
