package controllers

import (
	usecases "firstProject/internal/app/usecases/address"
	domain "firstProject/internal/domain/address"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	AddressUsecases *usecases.AddressUsecases
}

// NewAddressController krijon një instancë të AddressController
func NewAddressController(AddressUsecases *usecases.AddressUsecases) *AddressController {
	return &AddressController{
		AddressUsecases: AddressUsecases,
	}
}

// AddAddress godoc
// @Summary Add a new address
// @Description Add a new address to the database
// @Tags address
// @Accept json
// @Produce json
// @Param address body domain.Address true "New address details"
// @Success 201 {object} domain.Address "Address added successfully"
// @Router /address [post]
func (ac *AddressController) AddAddress(c *gin.Context) {
	var newAddress domain.Address
	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address data"})
		return
	}

	if err := newAddress.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAddress, err := ac.AddressUsecases.AddNewAddress(newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address"})
		return
	}

	c.JSON(http.StatusCreated, createdAddress)
}
