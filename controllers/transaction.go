package controllers

import (
	"flight-ticket-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var txn models.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var booking models.Booking
	if err := models.DB.First(&booking, txn.BookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	txn.Status = "paid" // default status
	models.DB.Create(&txn)
	c.JSON(http.StatusOK, txn)
}

func GetAllTransactions(c *gin.Context) {
	var txns []models.Transaction
	models.DB.Find(&txns)
	c.JSON(http.StatusOK, txns)
}
