package address

import (
	"errors"
	"time"
)

type Address struct {
	ID        uint      `json:"id"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *Address) Validate() error {
	if a.Street == "" {
		return errors.New("street is required")
	}
	if a.City == "" {
		return errors.New("city is required")
	}
	if a.State == "" {
		return errors.New("state is required")
	}
	if a.Country == "" {
		return errors.New("country is required")
	}
	if a.UserID == 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (Address) TableName() string {
	return "address"
}
