package models

type Address struct {
	ID      uint `gorm:"primaryKey"`
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

func GetAddresses() ([]Address, error) {
	var addresses []Address
	if err := db.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func CreateAddress(address Address) error {
	if err := db.Create(&address).Error; err != nil {
		return err
	}
	return nil
}
