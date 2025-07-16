package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	BookingID uint   `json:"booking_id"` // Foreign key
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}
