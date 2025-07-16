package controllers

import (
	"flight-ticket-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new flight
func CreateFlight(c *gin.Context) {
	var flight models.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&flight)
	c.JSON(http.StatusOK, flight)
}

// Get all flights
func GetAllFlights(c *gin.Context) {
	var flights []models.Flight
	models.DB.Find(&flights)
	c.JSON(http.StatusOK, flights)
}

// Get flight by ID
func GetFlightByID(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	if err := models.DB.First(&flight, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}
	c.JSON(http.StatusOK, flight)
}

// Update flight
func UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	if err := models.DB.First(&flight, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&flight)
	c.JSON(http.StatusOK, flight)
}

// Delete flight
func DeleteFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	if err := models.DB.First(&flight, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}
	models.DB.Delete(&flight)
	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted"})
}
