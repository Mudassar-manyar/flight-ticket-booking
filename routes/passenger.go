package routes

import (
	"flight-ticket-booking/controllers"

	"github.com/gin-gonic/gin"
)

func PassengerRoutes(r *gin.Engine) {
	passenger := r.Group("/passengers")
	{
		passenger.POST("/", controllers.CreatePassenger)
		passenger.GET("/", controllers.GetAllPassengers)
	}
}
