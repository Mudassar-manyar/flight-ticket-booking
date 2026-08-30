package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	FlightID uint   `json:"flight_id"`
	Status   string `json:"status"` // "booked", "cancelled", etc.
}
