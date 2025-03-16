package db

import (
	"firstProject/internal/domain/address"
	"fmt"

	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (r *AddressRepository) CreateAddress(newAddress *address.Address) (address.Address, error) {
	if err := r.DB.Create(newAddress).Error; err != nil {
		fmt.Println("Error while inserting address:", err)
		return address.Address{}, err
	}
	return *newAddress, nil
}
