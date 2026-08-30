package controllers

import (
	"flight-ticket-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new booking
func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: check if seats are available
	var flight models.Flight
	if err := models.DB.First(&flight, booking.FlightID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	if flight.SeatsAvailable <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No seats available"})
		return
	}

	// Decrease seat count
	flight.SeatsAvailable -= 1
	models.DB.Save(&flight)

	booking.Status = "confirmed"
	models.DB.Create(&booking)

	c.JSON(http.StatusOK, booking)
}

// Get all bookings
func GetAllBookings(c *gin.Context) {
	var bookings []models.Booking
	models.DB.Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}

// Get booking by ID
func GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking
	if err := models.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(http.StatusOK, booking)
}
