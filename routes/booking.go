package routes

import (
	"flight-ticket-booking/controllers"
	"flight-ticket-booking/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	booking := r.Group("/bookings")

	booking.Use(middleware.JWTAuthMiddleware()) //Protect all bookings , user + admin done
	{
		booking.POST("/", controllers.CreateBooking)
		booking.GET("/", controllers.GetAllBookings)
		booking.GET("/:id", controllers.GetBookingByID)
	}
}
