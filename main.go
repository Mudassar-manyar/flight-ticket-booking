package main

import (
	"flight-ticket-booking/config"
	"flight-ticket-booking/models"
	"flight-ticket-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to Database
	config.ConnectDatabase()

	// Auto Migrate Models
	config.DB.AutoMigrate(
		&models.User{},
		&models.Flight{},
		&models.Booking{},
		&models.Passenger{},
		&models.Transaction{},
	)

	//Home route only once
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Flight Booking API working!"})
	})

	//Register Routes
	routes.AuthRoutes(r)
	routes.FlightRoutes(r)
	routes.BookingRoutes(r)
	routes.PassengerRoutes(r)
	routes.TransactionRoutes(r)
	// Run server
	r.Run(":8080")
}
