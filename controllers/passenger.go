package controllers

import (
	"flight-ticket-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add passenger to booking
func CreatePassenger(c *gin.Context) {
	var passenger models.Passenger
	if err := c.ShouldBindJSON(&passenger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: validate booking exists
	var booking models.Booking
	if err := models.DB.First(&booking, passenger.BookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	models.DB.Create(&passenger)
	c.JSON(http.StatusOK, passenger)
}

// Get all passengers
func GetAllPassengers(c *gin.Context) {
	var passengers []models.Passenger
	models.DB.Find(&passengers)
	c.JSON(http.StatusOK, passengers)
}
