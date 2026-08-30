package models

import "gorm.io/gorm"

var DB *gorm.DB

type Flight struct {
	gorm.Model
	Airline        string  `json:"airline"`
	From           string  `json:"from"`
	To             string  `json:"to"`
	Departure      string  `json:"departure"` // use string for now (time.Time can be added later)
	Arrival        string  `json:"arrival"`
	Price          float64 `json:"price"`
	SeatsAvailable int     `json:"seats_available"`
}
