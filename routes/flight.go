package routes

import (
	"flight-ticket-booking/controllers"
	"flight-ticket-booking/middleware"

	"github.com/gin-gonic/gin"
)

func FlightRoutes(r *gin.Engine) {
	flight := r.Group("/flights")

	flight.Use(middleware.JWTAuthMiddleware(), middleware.AdminOnly()) //only admin
	{
		//flight.Use(middleware.JWTAuthMiddleware(), middleware.AdminOnly()) //only admin

		flight.POST("/", controllers.CreateFlight)
		flight.GET("/", controllers.GetAllFlights) // public
		flight.GET("/:id", controllers.GetFlightByID)
		flight.PUT("/:id", controllers.UpdateFlight)
		flight.DELETE("/:id", controllers.DeleteFlight)
	}
}
