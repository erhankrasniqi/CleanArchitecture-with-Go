package usecases

import (
	"firstProject/internal/domain/address"
	db "firstProject/internal/infrastructure/repository"
	"time"
)

type AddressUsecases struct {
	AddressRepo *db.AddressRepository
}

func NewAddressUsecases(repo *db.AddressRepository) *AddressUsecases {
	return &AddressUsecases{
		AddressRepo: repo,
	}
}

func (u *AddressUsecases) AddNewAddress(newAddress address.Address) (address.Address, error) {
	address := address.Address{
		Street:    newAddress.Street,
		City:      newAddress.City,
		State:     newAddress.State,
		Country:   newAddress.Country,
		UserID:    newAddress.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.AddressRepo.CreateAddress(&address)
}
