package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	BookingID   uint    `json:"booking_id"` // Foreign key
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`       // e.g. "paid", "pending", "failed"
	PaymentMode string  `json:"payment_mode"` // e.g. "card", "upi"
}
